package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/awesome-gocui/gocui"
	"github.com/jdpedrie/llmrpg/pkg/postgres"
	"github.com/urfave/cli/v2"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Println(fmt.Errorf("err: %w", err))
	}
}

func run(ctx context.Context) error {
	db, err := postgres.NewFromEnv()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	g, err := gocui.NewGui(gocui.OutputNormal, true)
	if err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	app := &cli.App{
		Name:     "llmrpg",
		Commands: []*cli.Command{},
	}

	return app.RunContext(ctx, os.Args)
}