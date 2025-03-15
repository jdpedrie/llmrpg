package game

import (
	"context"

	"github.com/geldata/gel-go"
	"github.com/jdpedrie/llmrpg/model"
)

type Manager struct {
	db *gel.Client
}

func NewManager(db *gel.Client) *Manager {
	return &Manager{db}
}

func (m *Manager) Templates(ctx context.Context) ([]model.Game, error) {
	games := make([]model.Game, 0)
	if err := m.db.Query(ctx, `SELECT Game FILTER .is_template = <bool> true`, &games); err != nil {
		return nil, err
	}

	return games, nil
}
