package game

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jdpedrie/llmrpg/model"
	"github.com/jdpedrie/llmrpg/pkg/postgres"
)

// Engine handles game mechanics and state management
type Engine struct {
	pool      *pgxpool.Pool
	queries   *postgres.Queries
	llmClient LLMClient
	gm        *GameMaster
}

// NewEngine creates a new game engine
func NewEngine(pool *pgxpool.Pool, queries *postgres.Queries, llmClient LLMClient) *Engine {
	engine := &Engine{
		pool:      pool,
		queries:   queries,
		llmClient: llmClient,
	}

	engine.gm = NewGameMaster(engine, llmClient)
	return engine
}

// CreateGame creates a new game template
func (e *Engine) CreateGame(ctx context.Context, game *model.Game) error {
	conn, err := e.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := e.queries.WithTx(tx)

	// Create the game
	gameParams := postgres.CreateGameParams{
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
	}

	pgGame, err := qtx.CreateGame(ctx, gameParams)
	if err != nil {
		return err
	}

	game.ID = pgGame.ID

	// Create characters
	for i := range game.Characters {
		char := &game.Characters[i]
		charParams := postgres.CreateCharacterParams{
			GameID:        game.ID,
			Name:          char.Name,
			Description:   postgres.NewText(char.Description),
			Context:       char.Context,
			Active:        char.Active,
			MainCharacter: char.MainCharacter,
		}

		pgCharacter, err := qtx.CreateCharacter(ctx, charParams)
		if err != nil {
			return err
		}

		char.ID = pgCharacter.ID

		for k, v := range map[string][]model.CharacterAttribute{
			"skill":          char.Skills,
			"characteristic": char.Characteristics,
			"relationship":   char.Relationship,
		} {
			// Create character attributes
			for _, attr := range v {
				attrParams := postgres.CreateCharacterAttributeParams{
					CharacterID:   pgCharacter.ID,
					AttributeType: k,
					Name:          attr.Name,
					Value:         attr.Value,
				}

				_, err := qtx.CreateCharacterAttribute(ctx, attrParams)
				if err != nil {
					return err
				}
			}
		}
	}

	// Create inventory items
	for i := range game.Inventory {
		item := &game.Inventory[i]
		itemParams := postgres.CreateInventoryItemParams{
			GameID:      game.ID,
			Name:        item.Name,
			Description: item.Description,
			Active:      item.Active,
		}

		pgItem, err := qtx.CreateInventoryItem(ctx, itemParams)
		if err != nil {
			return err
		}

		item.ID = pgItem.ID
	}

	return tx.Commit(ctx)
}

// GetGame returns a game by ID
func (e *Engine) GetGame(ctx context.Context, gameID string) (*model.Game, error) {
	gID, err := uuid.Parse(gameID)
	if err != nil {
		return nil, fmt.Errorf("invalid game ID: %w", err)
	}

	conn, err := e.pool.Acquire(ctx)
	if err != nil {
		return nil, err
	}

	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return nil, err
	}
	defer tx.Rollback(ctx)

	qtx := e.queries.WithTx(tx)

	// Get the game
	dbGame, err := qtx.GetGame(ctx, gID)
	if err != nil {
		return nil, err
	}

	// Get characters
	dbCharacters, err := qtx.GetGameCharacters(ctx, gID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	characters := make([]model.Character, 0, len(dbCharacters))
	for _, dbChar := range dbCharacters {
		// Get character attributes
		skills, err := qtx.GetCharacterAttributesByType(ctx, postgres.GetCharacterAttributesByTypeParams{
			CharacterID:   dbChar.ID,
			AttributeType: "skill",
		})
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		characteristics, err := qtx.GetCharacterAttributesByType(ctx, postgres.GetCharacterAttributesByTypeParams{
			CharacterID:   dbChar.ID,
			AttributeType: "characteristic",
		})
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		relationships, err := qtx.GetCharacterAttributesByType(ctx, postgres.GetCharacterAttributesByTypeParams{
			CharacterID:   dbChar.ID,
			AttributeType: "relationship",
		})
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, err
		}

		character := model.FromDBCharacter(dbChar, skills, characteristics, relationships)
		characters = append(characters, character)
	}

	// Get inventory items
	dbItems, err := qtx.GetGameInventory(ctx, gID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, err
	}

	items := make([]model.InventoryItem, 0, len(dbItems))
	for _, dbItem := range dbItems {
		items = append(items, model.FromDBInventoryItem(dbItem))
	}

	// Create the model game
	game := model.FromDBGame(dbGame, characters, items)

	if err := tx.Commit(ctx); err != nil {
		return nil, err
	}

	return &game, nil
}

// ReceiveAction processes a player action and returns streaming results
func (e *Engine) ReceiveAction(ctx context.Context, gameID string, action *model.Action, resultChan chan<- *model.ActionResult) error {
	defer close(resultChan)

	gID, err := uuid.Parse(gameID)
	if err != nil {
		return fmt.Errorf("invalid game ID: %w", err)
	}

	conn, err := e.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := e.queries.WithTx(tx)

	// Process the action through the game master
	responseChan := make(chan string)
	finalChan := make(chan bool)
	errorChan := make(chan error, 1)

	go func() {
		err := e.gm.ProcessAction(ctx, qtx, gID, action.Choice, action.Outcome, responseChan, finalChan)
		errorChan <- err
	}()

	// Stream responses back to the caller
	for {
		select {
		case message, ok := <-responseChan:
			if !ok {
				// Channel closed
				return nil
			}

			final := <-finalChan
			resultChan <- &model.ActionResult{
				Message: message,
				Final:   final,
			}

		case err := <-errorChan:
			if err != nil {
				return err
			}

			if err := tx.Commit(ctx); err != nil {
				return err
			}

			return nil

		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
