package server

import (
	"context"
	"errors"

	"connectrpc.com/connect"
	"github.com/jdpedrie/llmrpg/game"
	v1 "github.com/jdpedrie/llmrpg/genproto/jdpedrie/llmrpg/v1"
	"github.com/jdpedrie/llmrpg/genproto/jdpedrie/llmrpg/v1/v1connect"
	"github.com/jdpedrie/llmrpg/model"
	"github.com/jdpedrie/llmrpg/pkg/db"
)

type LLMRPCService struct {
	v1connect.UnimplementedLLMRPGServiceHandler

	manager *game.Manager
	engine  *game.Engine
}

func (s *LLMRPCService) StartGame(
	ctx context.Context, req *connect.Request[v1.StartGameRequest],
) (*connect.Response[v1.StartGameResponse], error) {
	game, err := s.engine.StartGame(ctx, req.Msg.TemplateId)
	if err != nil {
		return nil, err
	}

	return &connect.Response[v1.StartGameResponse]{
		Msg: &v1.StartGameResponse{
			GameId: game.GameID(),
		},
	}, nil
}

func (s *LLMRPCService) CreateGame(
	ctx context.Context, req *connect.Request[v1.CreateGameRequest],
) (*connect.Response[v1.CreateGameResponse], error) {
	var game model.Game
	if err := game.FromProto(req.Msg.Game); err != nil {
		return nil, err
	}

	if !db.UUIDEmpty(game.ID) {
		return nil, errors.New("game id must be empty")
	}

	if err := s.engine.CreateGame(ctx, &game); err != nil {
		return nil, err
	}

	// 2nd query because nested IDs not returned above.
	g, err := s.engine.GetGame(ctx, game.ID.String())
	if err != nil {
		return nil, err
	}

	return &connect.Response[v1.CreateGameResponse]{
		Msg: &v1.CreateGameResponse{
			Game: g.ToProto(),
		},
	}, nil
}
