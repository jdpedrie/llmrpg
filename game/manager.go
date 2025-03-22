package game

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jdpedrie/llmrpg/model"
	"github.com/jdpedrie/llmrpg/pkg/postgres"
)

// Manager handles game session lifecycle
type Manager struct {
	pool    *pgxpool.Pool
	queries *postgres.Queries
	engine  *Engine
}

// NewManager creates a new game manager
func NewManager(pool *pgxpool.Pool, queries *postgres.Queries) *Manager {
	return &Manager{
		pool:    pool,
		queries: queries,
	}
}

// SetEngine sets the engine for the manager
func (m *Manager) SetEngine(engine *Engine) {
	m.engine = engine
}

// ListGameTemplates lists all game templates
func (m *Manager) ListGameTemplates(ctx context.Context) ([]model.Game, error) {
	templates, err := m.queries.ListGameTemplates(ctx)
	if err != nil {
		return nil, err
	}

	games := make([]model.Game, 0, len(templates))
	for _, template := range templates {
		// Convert to model game without loading characters and inventory
		// This is for listing purposes only
		game := model.FromDBGame(template, nil, nil)
		games = append(games, game)
	}

	return games, nil
}

// ListActiveGames lists all active games
func (m *Manager) ListActiveGames(ctx context.Context) ([]model.Game, error) {
	activeGames, err := m.queries.ListActiveGames(ctx, postgres.ListActiveGamesParams{})
	if err != nil {
		return nil, err
	}

	games := make([]model.Game, 0, len(activeGames))
	for _, g := range activeGames {
		// Convert to model game without loading characters and inventory
		// This is for listing purposes only
		game := model.FromDBGame(g, nil, nil)
		games = append(games, game)
	}

	return games, nil
}

// EndGame marks a game as completed
func (m *Manager) EndGame(ctx context.Context, gameID string) error {
	gID, err := uuid.Parse(gameID)
	if err != nil {
		return err
	}

	conn, err := m.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := m.queries.WithTx(tx)

	if _, err := qtx.EndGame(ctx, gID); err != nil {
		return err
	}

	return tx.Commit(ctx)
}

// DeleteGame deletes a game and all related data
func (m *Manager) DeleteGame(ctx context.Context, gameID string) error {
	gID, err := uuid.Parse(gameID)
	if err != nil {
		return err
	}

	conn, err := m.pool.Acquire(ctx)
	if err != nil {
		return err
	}

	defer conn.Release()

	tx, err := conn.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := m.queries.WithTx(tx)

	// Check if the game exists
	_, err = qtx.GetGame(ctx, gID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("game not found")
		}
		return err
	}

	// Delete the game and all related data
	if err := qtx.DeleteGame(ctx, gID); err != nil {
		return err
	}

	return tx.Commit(ctx)
}
