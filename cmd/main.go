package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/jdpedrie/llmrpg/cmd/server"
	"github.com/jdpedrie/llmrpg/game"
	"github.com/jdpedrie/llmrpg/pkg/postgres"
	"github.com/urfave/cli/v2"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	if err := run(context.Background(), logger); err != nil {
		log.Println(fmt.Errorf("err: %w", err))
	}
}

func run(ctx context.Context, logger *slog.Logger) error {
	db, err := postgres.NewFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	queries := postgres.New(db)

	// Initialize game components
	manager := game.NewManager(db, queries)
	engine := game.NewEngine(db, queries, nil) // We'll need to implement an LLM client
	manager.SetEngine(engine)

	app := cli.App{
		Commands: []*cli.Command{
			{
				Name:   "serve",
				Action: server.Serve(logger.With("service", "server"), manager, engine),
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "port",
						Usage:   "change the server port",
						Aliases: []string{"p"},
						Value:   47788,
					},
				},
			},
		},
	}

	return app.RunContext(ctx, os.Args)
}
