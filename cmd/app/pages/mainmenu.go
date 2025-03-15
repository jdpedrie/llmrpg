package pages

import (
	fmt "fmt"

	tea "github.com/charmbracelet/bubbletea"
)

// We can define a custom message type to indicate
// when we want to switch pages or quit.
type SwitchPageMsg struct {
	PageName string
}

type MainMenuModel struct {
	cursor    int
	menuItems []string
}

func NewMainMenuModel() *MainMenuModel {
	return &MainMenuModel{
		menuItems: []string{"New Game", "Open Screen", "Wizard Form", "Quit"},
	}
}

// Satisfy the Page interface.
func (m *MainMenuModel) Init() tea.Cmd {
	return nil
}

func (m *MainMenuModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down":
			if m.cursor < len(m.menuItems)-1 {
				m.cursor++
			}
		case "enter":
			switch m.cursor {
			case 0:
				return m, func() tea.Msg {
					return SwitchPageMsg{PageName: "newgame"}
				}
				// // Switch to Screen
				// return m, func() tea.Msg {
				// 	return SwitchPageMsg{PageName: "screen"}
				// }
			case 1:
				// Switch to Wizard
				return m, func() tea.Msg {
					return SwitchPageMsg{PageName: "wizard"}
				}
			case 2:
				// Quit
				return m, tea.Quit
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m *MainMenuModel) View() string {
	s := "Welcome to llmrpg\n\n"
	for i, item := range m.menuItems {
		cursor := "  "
		if i == m.cursor {
			cursor = "> "
		}
		s += fmt.Sprintf("%s%s\n", cursor, item)
	}
	s += "\nPress 'q' or Ctrl+C to quit."
	return s
}
