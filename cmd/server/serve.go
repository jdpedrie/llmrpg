package server

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"connectrpc.com/grpcreflect"
	"github.com/jdpedrie/llmrpg/game"
	"github.com/jdpedrie/llmrpg/genproto/jdpedrie/llmrpg/v1/v1connect"
	"github.com/urfave/cli/v2"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

func Serve(logger *slog.Logger, manager *game.Manager, engine *game.Engine) cli.ActionFunc {
	return func(ctx *cli.Context) error {
		logger = logger.With("port", ctx.Int("port"))

		svc := LLMRPCService{
			manager: manager,
			engine:  engine,
		}

		mux := http.NewServeMux()

		r := grpcreflect.NewStaticReflector(v1connect.LLMRPGServiceName)
		mux.Handle(v1connect.NewLLMRPGServiceHandler(&svc))
		mux.Handle(grpcreflect.NewHandlerV1(r))
		mux.Handle(grpcreflect.NewHandlerV1Alpha(r))

		server := &http.Server{
			Addr:    fmt.Sprintf(":%d", ctx.Int("port")),
			Handler: h2c.NewHandler(mux, &http2.Server{}),
		}

		// Listen for shutdown signals
		go func() {
			<-ctx.Context.Done()
			logger.Info("server stopping...")
			_ = server.Shutdown(context.Background())
		}()

		logger.Info("server started")

		return server.ListenAndServe()
	}
}
