package game

import (
	"context"

	"github.com/geldata/gel-go"
	"github.com/geldata/gel-go/geltypes"
	"github.com/jdpedrie/llmrpg/model"
	"github.com/jdpedrie/llmrpg/pkg/db"
)

type Engine struct {
	db *gel.Client
}

func NewEngine(db *gel.Client) *Engine {
	return &Engine{
		db: db,
	}
}

func (e *Engine) StartGame(ctx context.Context, templateID string) (*GameEngine, error) {
	templateUUID, err := geltypes.ParseUUID(templateID)
	if err != nil {
		return nil, err
	}

	var tpl model.Game
	q := `SELECT Game {**} FILTER .id = <uuid>$0 AND .is_template = <bool>TRUE`
	if err := e.db.QuerySingle(ctx, q, &tpl, templateUUID); err != nil {
		return nil, err
	}

	return &GameEngine{
		db: e.db,
		// gameID: gameID,
	}, nil
}

func (e *Engine) CreateGame(ctx context.Context, game *model.Game) error {
	return db.Insert(ctx, e.db, game)
}

type GameEngine struct {
	db     *gel.Client
	gameID string
}

func (g *GameEngine) GameID() string {
	return g.gameID
}

func (e *Engine) CurrentState(ctx context.Context) (any, error) {
	return nil, nil
}

func (e *Engine) Action(ctx context.Context, action any) (chan<- any, error) {
	return nil, nil
}
