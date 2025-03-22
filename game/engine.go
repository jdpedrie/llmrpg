package game

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jdpedrie/llmrpg/model"
	"github.com/jdpedrie/llmrpg/pkg/postgres"
)

// Engine manages game engine functionality
type Engine struct {
	db *postgres.DB
}

// NewEngine creates a new game engine
func NewEngine(db *postgres.DB) *Engine {
	return &Engine{
		db: db,
	}
}

// StartGame creates a new game instance from a template
func (e *Engine) StartGame(ctx context.Context, templateID string) (*GameEngine, error) {
	gameID, err := uuid.Parse(templateID)
	if err != nil {
		return nil, fmt.Errorf("invalid template ID: %w", err)
	}

	// Retrieve template game
	var tpl *model.Game
	manager := NewManager(e.db)
	tpl, err = manager.GetGame(ctx, gameID.String())
	if err != nil {
		return nil, fmt.Errorf("error getting template game: %w", err)
	}

	if !tpl.IsTemplate {
		return nil, fmt.Errorf("game %s is not a template", templateID)
	}

	// Make a copy of the template with IsTemplate=false and IsRunning=true
	tpl.ID = uuid.New() // Generate a new ID for the game instance
	tpl.IsTemplate = false
	tpl.IsRunning = true

	// Reset IDs for all entities to ensure new instances are created
	for i := range tpl.Characters {
		tpl.Characters[i].ID = uuid.New()
		
		for j := range tpl.Characters[i].Skills {
			tpl.Characters[i].Skills[j].ID = uuid.New()
		}
		
		for j := range tpl.Characters[i].Characteristics {
			tpl.Characters[i].Characteristics[j].ID = uuid.New()
		}
		
		for j := range tpl.Characters[i].Relationship {
			tpl.Characters[i].Relationship[j].ID = uuid.New()
		}
	}

	for i := range tpl.Inventory {
		tpl.Inventory[i].ID = uuid.New()
	}

	// Use a transaction for creating the new game
	err = e.db.WithTx(ctx, func(q *postgres.Queries) error {
		// Create the game first
		dbGame, err := q.CreateGame(ctx, postgres.CreateGameParams{
			Name:            tpl.Name,
			Description:     postgres.NewText(tpl.Description),
			StartingMessage: postgres.NewText(tpl.StartingMessage),
			Scenario:        postgres.NewText(tpl.Scenario),
			Objectives:      postgres.NewText(tpl.Objectives),
			Skills:          tpl.Skills,
			Characteristics: tpl.Characteristics,
			Relationship:    tpl.Relationship,
			IsTemplate:      false,
			IsRunning:       true,
		})
		if err != nil {
			return fmt.Errorf("error creating game from template: %w", err)
		}

		tpl.ID = dbGame.ID
		
		// Create characters and attributes
		for i := range tpl.Characters {
			char := &tpl.Characters[i]
			dbChar, err := q.CreateCharacter(ctx, postgres.CreateCharacterParams{
				Name:          char.Name,
				Description:   postgres.NewText(char.Description),
				Context:       char.Context,
				Active:        char.Active,
				MainCharacter: char.MainCharacter,
				GameID:        gameID{ID: tpl.ID, Valid: true},
			})
			if err != nil {
				return fmt.Errorf("error creating character: %w", err)
			}

			char.ID = dbChar.ID

			// Create character attributes
			if err := createCharacterAttributes(ctx, q, char); err != nil {
				return err
			}
		}

		// Create inventory items
		for i := range tpl.Inventory {
			item := &tpl.Inventory[i]
			dbItem, err := q.CreateInventoryItem(ctx, postgres.CreateInventoryItemParams{
				Name:        item.Name,
				Description: item.Description,
				Active:      item.Active,
				GameID:      tpl.ID,
			})
			if err != nil {
				return fmt.Errorf("error creating inventory item: %w", err)
			}

			item.ID = dbItem.ID
			item.GameID = dbItem.GameID
		}

		// Update the game to set the playthrough_start_time
		_, err = q.StartGame(ctx, tpl.ID)
		if err != nil {
			return fmt.Errorf("error starting game: %w", err)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return &GameEngine{
		db:     e.db,
		gameID: tpl.ID.String(),
	}, nil
}

// CreateGame creates a new game
func (e *Engine) CreateGame(ctx context.Context, game *model.Game) error {
	manager := NewManager(e.db)
	return manager.CreateGame(ctx, game)
}

// GetGame retrieves a game by ID
func (e *Engine) GetGame(ctx context.Context, id string) (*model.Game, error) {
	manager := NewManager(e.db)
	return manager.GetGame(ctx, id)
}

// GameEngine represents an active game instance
type GameEngine struct {
	db     *postgres.DB
	gameID string
}

// GameID returns the ID of the active game
func (g *GameEngine) GameID() string {
	return g.gameID
}

// CurrentState returns the current state of the game
func (e *Engine) CurrentState(ctx context.Context) (any, error) {
	return nil, nil
}

// Action processes a player action
func (e *GameEngine) Action(ctx context.Context, action PlayerAction) (<-chan any, error) {
	return nil, nil
}

// PlayerAction represents a player's action in the game
type PlayerAction struct {
	Choice  string
	Success string
}

// ActionResult represents the result of a player action
type ActionResult struct {
	Message string
	Game    *model.Game
}