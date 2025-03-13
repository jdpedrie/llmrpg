package pages

import (
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
)

type ScreenModel struct {
	width  int
	height int
	// any other fields you need
}

func NewScreenModel() *ScreenModel {
	return &ScreenModel{}
}

// Init: Enter the alternate screen here
func (m *ScreenModel) Init() tea.Cmd {
	return tea.Batch(func() tea.Msg {
		// If you’re on a Unix-like system, you can do something like this:
		w, h, err := term.GetSize(int(os.Stdin.Fd()))
		if err != nil {
			// Fallback if something goes wrong
			w, h = 80, 24
		}
		return tea.WindowSizeMsg{Width: w, Height: h}
	}, tea.EnterAltScreen)
}

// Update handles events.
func (m *ScreenModel) Update(msg tea.Msg) (Page, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc", "enter":
			// We want to EXIT the alt screen before returning to main menu.
			// So we chain two commands: ExitAltScreen, then SwitchPageMsg.
			return m, tea.Sequence(
				tea.ExitAltScreen,
				func() tea.Msg {
					return SwitchPageMsg{PageName: "mainmenu"}
				},
			)
		}
	}
	return m, nil
}

// View draws the full-screen layout via Lip Gloss.
func (m *ScreenModel) View() string {
	// If we haven't received a WindowSizeMsg yet, just display a placeholder.
	if m.width == 0 || m.height == 0 {
		return "Loading full-screen layout..."
	}

	// Example layout with a sidebar, main area, and footer:
	sidebarStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2)

	mainStyle := lipgloss.NewStyle().
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("63")).
		Padding(1, 2)

	footerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("241")).
		Background(lipgloss.Color("236")).
		Padding(0, 1)

	sidebarWidth := 20
	if sidebarWidth > m.width-10 {
		sidebarWidth = m.width / 3
	}
	mainWidth := m.width - sidebarWidth
	usableHeight := m.height - 2
	if usableHeight < 1 {
		usableHeight = 1
	}

	sidebarBox := sidebarStyle.
		Width(sidebarWidth).
		Height(usableHeight).
		Render("SIDEBAR\n\n- Item 1\n- Item 2\n- Item 3")

	mainBox := mainStyle.
		Width(mainWidth).
		Height(usableHeight).
		Render("MAIN CONTENT\n\nPress Q/Esc/Enter to return to the main menu.")

	body := lipgloss.JoinHorizontal(lipgloss.Top, sidebarBox, mainBox)

	footer := footerStyle.
		Width(m.width).
		Render("[Q/Esc/Enter] Return to Main Menu")

	return lipgloss.JoinVertical(lipgloss.Top, body, footer)
}
