package main

import (
	"context"
	"fmt"
	"log"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/geldata/gel-go"
	"github.com/geldata/gel-go/gelcfg"
	"github.com/jdpedrie/llmrpg/cmd/app/pages"
)

func main() {
	if err := run(context.Background()); err != nil {
		log.Println(fmt.Errorf("err: %w", err))
	}
}

func run(ctx context.Context) error {
	client, err := gel.CreateClient(gelcfg.Options{})
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	// Start with the Main Menu
	initialModel := model{
		currentPage: pages.NewMainMenuModel(),
	}

	p := tea.NewProgram(initialModel)
	_, err = p.Run()
	return err
}

type model struct {
	currentPage pages.Page
}

// Implement tea.Model for our top-level model.

func (m model) Init() tea.Cmd {
	// Initialize whatever the current page is
	return m.currentPage.Init()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	// First, let the current page handle the update
	newPage, cmd := m.currentPage.Update(msg)

	// If the page wants to switch to another page, we handle that
	switch msg := msg.(type) {
	case pages.SwitchPageMsg:
		switch msg.PageName {
		case "mainmenu":
			newPage = pages.NewMainMenuModel()
		case "screen":
			newPage = pages.NewScreenModel()
		case "wizard":
			newPage = pages.NewWizardModel()
		}
		// We call Init() on the new page in case it needs to set up
		initCmd := newPage.Init()
		return model{currentPage: newPage}, tea.Batch(cmd, initCmd)
	}

	// Otherwise, just continue with the same page
	return model{currentPage: newPage}, cmd
}

func (m model) View() string {
	return m.currentPage.View()
}
