package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	"github.com/awesome-gocui/gocui"
	"github.com/jdpedrie/llmrpg/cmd/server"
	"github.com/jdpedrie/llmrpg/game"
	"github.com/jdpedrie/llmrpg/pkg/postgres"
	"github.com/urfave/cli/v2"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Println(fmt.Errorf("err: %w", err))
	}
}

func run(ctx context.Context) error {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

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

	app := &cli.App{
		Name: "llmrpg",
		Commands: []*cli.Command{
			{
				Name:  "server",
				Usage: "Start the LLMRPG API server",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "port",
						Aliases: []string{"p"},
						Value:   8080,
						Usage:   "Server port",
					},
				},
				Action: server.Serve(logger, manager, engine),
			},
			{
				Name:  "gui",
				Usage: "Start the LLMRPG GUI client",
				Action: func(c *cli.Context) error {
					g, err := gocui.NewGui(gocui.OutputNormal, true)
					if err != nil {
						return err
					}
					defer g.Close()
					
					// We'll implement the GUI later
					return nil
				},
			},
		},
	}

	return app.RunContext(ctx, os.Args)
}