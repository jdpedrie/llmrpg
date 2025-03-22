package main

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"

	tea "github.com/charmbracelet/bubbletea"
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
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	defer db.Close()

	manager := game.NewManager(db)
	engineFactory := game.NewEngine(db)

	app := cli.App{
		Action: func(ctx *cli.Context) error {
			e, err := NewEntryModel(&page1{}, &page2{})
			if err != nil {
				return err
			}
			p := tea.NewProgram(e, tea.WithContext(ctx.Context))
			_, err = p.Run()

			return err
		},
		Commands: []*cli.Command{
			{
				Name:   "serve",
				Action: server.Serve(logger.With("service", "server"), manager, engineFactory),
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

type Page interface {
	tea.Model

	ID() string
}

type EntryModel struct {
	currentState string
	pages        map[string]Page
}

func NewEntryModel(entry Page, pages ...Page) (tea.Model, error) {
	p := make(map[string]Page)
	p[""] = entry
	for _, page := range pages {
		if _, ok := p[page.ID()]; ok {
			return nil, fmt.Errorf("id %s already registered", page.ID())
		}

		p[page.ID()] = page
	}

	return &EntryModel{
		pages: p,
	}, nil
}

func (m *EntryModel) View() string {
	p, err := m.getModel()
	if err != nil {
		return err.Error()
	}

	return p.View()
}

func (m *EntryModel) Init() tea.Cmd {
	p, err := m.getModel()
	if err != nil {
		return nil
	}

	return p.Init()
}

func (m *EntryModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	p, err := m.getModel()
	if err != nil {
		return m, nil
	}

	model, cmd := p.Update(msg)

	if cmd == nil {
		return model, nil
	}

	msg = cmd()
	switch x := msg.(type) {
	case changePage:
		m.currentState = x.id
	case tea.KeyMsg:
		if x.String() == "ctrl+c" {
			return model, tea.Quit
		}
	}

	return model, nil
}

func (m *EntryModel) getModel() (tea.Model, error) {
	page, ok := m.pages[m.currentState]
	if !ok {
		return nil, fmt.Errorf("invalid current state id [%s]", m.currentState)
	}

	return page, nil
}

type page1 struct{}

func (page1) ID() string {
	return "mainmenu"
}

func (m *page1) Init() tea.Cmd {
	return nil
}

func (m *page1) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch x := msg.(type) {
	case tea.KeyMsg:
		if x.Type == tea.KeyEnter {
			return m, ChangePage("page2")
		}
	}

	return m, nil
}

func (m *page1) View() string {
	return "press enter to continue"
}

type page2 struct{}

func (page2) ID() string {
	return "mainmenu"
}

func (m *page2) Init() tea.Cmd {
	return nil
}

func (m *page2) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch x := msg.(type) {
	case tea.KeyMsg:
		if x.Type == tea.KeyEsc {
			return m, ChangePage("page1")
		}
	}

	return m, nil
}

func (m *page2) View() string {
	return "press esc to return"
}

func ChangePage(id string) tea.Cmd {
	return func() tea.Msg {
		return changePage{
			id: id,
		}
	}
}

type changePage struct {
	id string
}