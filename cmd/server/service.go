package server

import (
	"context"
	"errors"
	"io"
	"log"

	"connectrpc.com/connect"
	"github.com/jdpedrie/llmrpg/game"
	v1 "github.com/jdpedrie/llmrpg/genproto/jdpedrie/llmrpg/v1"
	"github.com/jdpedrie/llmrpg/genproto/jdpedrie/llmrpg/v1/v1connect"
	"github.com/jdpedrie/llmrpg/model"
)

type LLMRPCService struct {
	v1connect.UnimplementedLLMRPGServiceHandler

	manager *game.Manager
	engine  *game.Engine
}

func NewLLMRPCService(manager *game.Manager, engine *game.Engine) *LLMRPCService {
	return &LLMRPCService{
		manager: manager,
		engine:  engine,
	}
}

func (s *LLMRPCService) Play(
	ctx context.Context, stream *connect.BidiStream[v1.PlayRequest, v1.PlayResponse],
) error {
	for {
		req, err := stream.Receive()
		if err != nil {
			if errors.Is(err, io.EOF) {
				return nil
			}
			return err
		}

		// Get the current game state
		gameID := req.GameId
		if gameID == "" {
			return connect.NewError(connect.CodeInvalidArgument, errors.New("game_id is required"))
		}

		game, err := s.engine.GetGame(ctx, gameID)
		if err != nil {
			return connect.NewError(connect.CodeNotFound, errors.New("game not found"))
		}

		// Send initial game state to client
		if err := stream.Send(&v1.PlayResponse{
			Resp: &v1.PlayResponse_Game{
				Game: game.ToProto(),
			},
		}); err != nil {
			return err
		}

		// Process user action if provided
		if req.Choice != "" {
			// Create a channel to receive streaming responses from the game engine
			responseChan := make(chan *model.ActionResult)
			errorChan := make(chan error, 1)

			// Process the action in a goroutine
			go func() {
				err := s.engine.ReceiveAction(ctx, gameID, &model.Action{
					Choice:  req.Choice,
					Outcome: req.Outcome,
				}, responseChan)
				errorChan <- err
			}()

			// Stream responses back to client
		processingLoop:
			for {
				select {
				case result, ok := <-responseChan:
					if !ok {
						// Channel closed, action processing complete
						break processingLoop
					}

					// Send the message to the client
					err := stream.Send(&v1.PlayResponse{
						Resp: &v1.PlayResponse_Message{
							Message: result.Message,
						},
					})
					if err != nil {
						log.Printf("Error sending response: %v", err)
						return err
					}

					// If this is the final response, send the updated game state
					if result.Final {
						updatedGame, err := s.engine.GetGame(ctx, gameID)
						if err != nil {
							return err
						}

						if err := stream.Send(&v1.PlayResponse{
							Resp: &v1.PlayResponse_Game{
								Game: updatedGame.ToProto(),
							},
						}); err != nil {
							return err
						}
					}
				case err := <-errorChan:
					if err != nil {
						return connect.NewError(connect.CodeInternal, err)
					}
					// Action completed successfully
					break processingLoop
				case <-ctx.Done():
					return ctx.Err()
				}
			}
		}
	}
}
