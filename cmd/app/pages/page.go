package pages

import (
	tea "github.com/charmbracelet/bubbletea"
)

// Page is the interface that all “pages” in our application will implement.
type Page interface {
	Init() tea.Cmd
	Update(msg tea.Msg) (Page, tea.Cmd)
	View() string
}
