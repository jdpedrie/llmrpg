package game

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jdpedrie/llmrpg/model"
	"github.com/jdpedrie/llmrpg/pkg/postgres"
)

// Manager handles game operations
type Manager struct {
	db *postgres.DB
}

// NewManager creates a new game manager
func NewManager(db *postgres.DB) *Manager {
	return &Manager{db: db}
}

// Templates returns all game templates
func (m *Manager) Templates(ctx context.Context) ([]model.Game, error) {
	templates, err := m.db.ListGameTemplates(ctx)
	if err != nil {
		return nil, err
	}

	games := make([]model.Game, 0, len(templates))
	for _, t := range templates {
		// For templates listing, we don't need to load the full game data with characters and inventory
		game := model.Game{
			ID:              t.ID,
			Name:            t.Name,
			Description:     postgres.StringFromText(t.Description),
			StartingMessage: postgres.StringFromText(t.StartingMessage),
			Scenario:        postgres.StringFromText(t.Scenario),
			Objectives:      postgres.StringFromText(t.Objectives),
			Skills:          t.Skills,
			Characteristics: t.Characteristics,
			Relationship:    t.Relationship,
			IsTemplate:      t.IsTemplate,
			IsRunning:       t.IsRunning,
			CreatedAt:       t.CreatedAt.Time,
			UpdatedAt:       t.UpdatedAt.Time,
		}
		games = append(games, game)
	}

	return games, nil
}

// GetGame retrieves a game by ID with all its characters and inventory
func (m *Manager) GetGame(ctx context.Context, id string) (*model.Game, error) {
	gameID, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid game ID: %w", err)
	}

	// We need to execute several queries to fetch the full game data
	var result *model.Game
	err = m.db.WithTx(ctx, func(q *postgres.Queries) error {
		// Get the game
		dbGame, err := q.GetGame(ctx, gameID)
		if err != nil {
			return err
		}

		// Get the characters
		dbCharacters, err := q.ListCharacters(ctx, gameID)
		if err != nil {
			return err
		}

		// Get the inventory
		dbInventory, err := q.ListGameInventory(ctx, gameID)
		if err != nil {
			return err
		}

		// Build the full game with characters and their attributes
		characters := make([]model.Character, 0, len(dbCharacters))
		for _, dbChar := range dbCharacters {
			// Get character skills
			skills, err := q.ListCharacterAttributesByType(ctx, postgres.ListCharacterAttributesByTypeParams{
				CharacterID:      dbChar.ID,
				RelationshipType: "skill",
			})
			if err != nil {
				return err
			}

			// Get character characteristics
			characteristics, err := q.ListCharacterAttributesByType(ctx, postgres.ListCharacterAttributesByTypeParams{
				CharacterID:      dbChar.ID,
				RelationshipType: "characteristic",
			})
			if err != nil {
				return err
			}

			// Get character relationships
			relationships, err := q.ListCharacterAttributesByType(ctx, postgres.ListCharacterAttributesByTypeParams{
				CharacterID:      dbChar.ID,
				RelationshipType: "relationship",
			})
			if err != nil {
				return err
			}

			character := model.FromDBCharacter(dbChar, skills, characteristics, relationships)
			characters = append(characters, character)
		}

		// Convert inventory items
		inventory := make([]model.InventoryItem, 0, len(dbInventory))
		for _, item := range dbInventory {
			inventory = append(inventory, model.FromDBInventoryItem(item))
		}

		// Create the full game model
		game := model.FromDBGame(dbGame, characters, inventory)
		result = &game
		return nil
	})

	if err != nil {
		return nil, err
	}

	return result, nil
}

// CreateGame creates a new game
func (m *Manager) CreateGame(ctx context.Context, game *model.Game) error {
	if !postgres.IsUUIDEmpty(game.ID) {
		return errors.New("game ID must be empty for creation")
	}

	// Use a transaction to create the game and all related entities
	return m.db.WithTx(ctx, func(q *postgres.Queries) error {
		// Create the game first
		dbGame, err := q.CreateGame(ctx, postgres.CreateGameParams{
			Name:            game.Name,
			Description:     postgres.NewText(game.Description),
			StartingMessage: postgres.NewText(game.StartingMessage),
			Scenario:        postgres.NewText(game.Scenario),
			Objectives:      postgres.NewText(game.Objectives),
			Skills:          game.Skills,
			Characteristics: game.Characteristics,
			Relationship:    game.Relationship,
			IsTemplate:      game.IsTemplate,
			IsRunning:       game.IsRunning,
		})
		if err != nil {
			return fmt.Errorf("error creating game: %w", err)
		}

		game.ID = dbGame.ID
		game.CreatedAt = dbGame.CreatedAt.Time
		game.UpdatedAt = dbGame.UpdatedAt.Time

		// Create all characters
		for i := range game.Characters {
			char := &game.Characters[i]
			dbChar, err := q.CreateCharacter(ctx, postgres.CreateCharacterParams{
				Name:          char.Name,
				Description:   postgres.NewText(char.Description),
				Context:       char.Context,
				Active:        char.Active,
				MainCharacter: char.MainCharacter,
				GameID:        gameID{ID: game.ID, Valid: true},
			})
			if err != nil {
				return fmt.Errorf("error creating character: %w", err)
			}

			char.ID = dbChar.ID
			char.CreatedAt = dbChar.CreatedAt.Time
			char.UpdatedAt = dbChar.UpdatedAt.Time

			// Create character attributes (skills, characteristics, relationships)
			if err := createCharacterAttributes(ctx, q, char); err != nil {
				return err
			}
		}

		// Create inventory items
		for i := range game.Inventory {
			item := &game.Inventory[i]
			dbItem, err := q.CreateInventoryItem(ctx, postgres.CreateInventoryItemParams{
				Name:        item.Name,
				Description: item.Description,
				Active:      item.Active,
				GameID:      game.ID,
			})
			if err != nil {
				return fmt.Errorf("error creating inventory item: %w", err)
			}

			item.ID = dbItem.ID
			item.GameID = dbItem.GameID
			item.CreatedAt = dbItem.CreatedAt.Time
			item.UpdatedAt = dbItem.UpdatedAt.Time
		}

		return nil
	})
}

// Helper function to create and link character attributes
func createCharacterAttributes(ctx context.Context, q *postgres.Queries, char *model.Character) error {
	// Create skills
	for i := range char.Skills {
		attr := &char.Skills[i]
		dbAttr, err := q.CreateCharacterAttribute(ctx, postgres.CreateCharacterAttributeParams{
			Name:          attr.Name,
			Value:         attr.Value,
			AttributeType: "skill",
		})
		if err != nil {
			return fmt.Errorf("error creating skill attribute: %w", err)
		}

		attr.ID = dbAttr.ID
		attr.Type = dbAttr.AttributeType

		// Link the attribute to the character
		err = q.LinkCharacterAttribute(ctx, postgres.LinkCharacterAttributeParams{
			CharacterID:      char.ID,
			AttributeID:      attr.ID,
			RelationshipType: "skill",
		})
		if err != nil {
			return fmt.Errorf("error linking skill attribute: %w", err)
		}
	}

	// Create characteristics
	for i := range char.Characteristics {
		attr := &char.Characteristics[i]
		dbAttr, err := q.CreateCharacterAttribute(ctx, postgres.CreateCharacterAttributeParams{
			Name:          attr.Name,
			Value:         attr.Value,
			AttributeType: "characteristic",
		})
		if err != nil {
			return fmt.Errorf("error creating characteristic attribute: %w", err)
		}

		attr.ID = dbAttr.ID
		attr.Type = dbAttr.AttributeType

		// Link the attribute to the character
		err = q.LinkCharacterAttribute(ctx, postgres.LinkCharacterAttributeParams{
			CharacterID:      char.ID,
			AttributeID:      attr.ID,
			RelationshipType: "characteristic",
		})
		if err != nil {
			return fmt.Errorf("error linking characteristic attribute: %w", err)
		}
	}

	// Create relationships
	for i := range char.Relationship {
		attr := &char.Relationship[i]
		dbAttr, err := q.CreateCharacterAttribute(ctx, postgres.CreateCharacterAttributeParams{
			Name:          attr.Name,
			Value:         attr.Value,
			AttributeType: "relationship",
		})
		if err != nil {
			return fmt.Errorf("error creating relationship attribute: %w", err)
		}

		attr.ID = dbAttr.ID
		attr.Type = dbAttr.AttributeType

		// Link the attribute to the character
		err = q.LinkCharacterAttribute(ctx, postgres.LinkCharacterAttributeParams{
			CharacterID:      char.ID,
			AttributeID:      attr.ID,
			RelationshipType: "relationship",
		})
		if err != nil {
			return fmt.Errorf("error linking relationship attribute: %w", err)
		}
	}

	return nil
}

// gameID is a helper type for handling nullable UUIDs
type gameID struct {
	ID    uuid.UUID
	Valid bool
}
