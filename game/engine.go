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
	q := `SELECT Game {
		*,
		characters: {
			*,
			skills: {**},
			characteristics: {**},
			relationship: {**}
		}
	} FILTER .id = <uuid>$0 AND .is_template = <bool>TRUE`
	if err := e.db.QuerySingle(ctx, q, &tpl, templateUUID); err != nil {
		return nil, err
	}

	tpl.IsTemplate = false
	tpl.IsRunning = true

	if err := e.CreateGame(ctx, &tpl); err != nil {
		return nil, err
	}

	return &GameEngine{
		db:     e.db,
		gameID: tpl.ID.String(),
	}, nil
}

func (e *Engine) CreateGame(ctx context.Context, game *model.Game) error {
	return db.Insert(ctx, e.db, game)
}

func (e *Engine) GetGame(ctx context.Context, id string) (*model.Game, error) {
	gameID, err := geltypes.ParseUUID(id)
	if err != nil {
		return nil, err
	}

	var game model.Game
	if err := e.db.QuerySingle(ctx, `SELECT Game{
		*, Character.**} FILTER .id = <uuid>$0`, &game, gameID); err != nil {
		return nil, err
	}

	return &game, nil
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
