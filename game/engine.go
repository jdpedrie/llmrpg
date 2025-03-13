package game

import (
	"context"

	"github.com/geldata/gel-go"
)

type Engine struct {
	db     *gel.Client
	gameID string
}

func NewEngine(db *gel.Client, gameID string) *Engine {
	return &Engine{
		db:     db,
		gameID: gameID,
	}
}

func (e *Engine) CurrentState(ctx context.Context) (any, error) {
	return nil, nil
}

func (e *Engine) Action(ctx context.Context, action any) (chan<- any, error) {
	return nil, nil
}
