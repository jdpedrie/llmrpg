package pages

import (
	"context"
	fmt "fmt"

	"github.com/charmbracelet/bubbles/table"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/jdpedrie/llmrpg/game"
)

type TemplatesListModel struct {
	table      table.Model
	lastAction string

	gameManager *game.Manager
}

// The action message we’ll emit when the user presses Enter on a row.
type performActionMsg struct {
	RowData table.Row
}

func NewTemplatesListModel(m *game.Manager) *TemplatesListModel {
	// 1) Define the columns in our table.
	columns := []table.Column{
		{Title: "Name", Width: 20},
	}

	// // 2) Define the rows of data.
	// // Each row is a slice of strings that match our column definitions.
	// rows := []table.Row{
	// 	{"1", "Alice", "Admin"},
	// 	{"2", "Bob", "Editor"},
	// 	{"3", "Charlie", "Viewer"},
	// 	{"4", "Diana", "Editor"},
	// 	{"5", "Ethan", "Admin"},
	// }

	// 3) Create a new table model.
	t := table.New(
		table.WithColumns(columns),
		// table.WithRows(rows),
		table.WithFocused(true),
	)

	// Optional: Set the initial visible rows or styles
	t.SetWidth(60)
	t.SetHeight(10)

	return &TemplatesListModel{
		table: t,

		gameManager: m,
	}
}

// Init is where we can perform initial commands (none needed here).
func (m TemplatesListModel) Init() tea.Cmd {
	return nil
}

func (m TemplatesListModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {

	case tea.KeyMsg:
		switch msg.String() {
		case "q", "esc":
			// We want to EXIT the alt screen before returning to main menu.
			// So we chain two commands: ExitAltScreen, then SwitchPageMsg.
			return m, tea.Sequence(
				tea.ExitAltScreen,
				func() tea.Msg {
					return SwitchPageMsg{PageName: "mainmenu"}
				},
			)

		// Pressing Enter triggers an action on the selected row
		case "enter":
			selectedRow := m.table.SelectedRow()
			return m, func() tea.Msg {
				return performActionMsg{RowData: selectedRow}
			}
		}

		// Let the table handle keypresses (arrows for navigation, etc.)
		newTable, tableCmd := m.table.Update(msg)
		m.table = newTable
		return m, tableCmd

	// Our custom message that indicates an action was performed.
	case performActionMsg:
		row := msg.RowData
		m.lastAction = fmt.Sprintf("Performed action on row: ID=%v, Name=%v, Role=%v",
			row[0], row[1], row[2])
		return m, nil
	}

	// If none of the above, pass the message to the table.
	m.table, cmd = m.table.Update(msg)
	return m, cmd
}

func (m TemplatesListModel) View() string {
	templates, err := m.gameManager.Templates(context.Background())
	if err != nil {
		return fmt.Sprintf("could not fetch game templates: %s", err.Error())
	}

	rows := make([]table.Row, 0, len(templates))
	for _, t := range templates {
		rows = append(rows, []string{
			t.Name,
		})
	}

	m.table.SetRows(rows)
	s := "Choose a game\n\n"
	s += m.table.View() + "\n"

	s += "\nNavigate with ↑/↓ (and ←/→ if your table is wide). Press Enter to take an action on the selected row.\n"
	s += "Press 'q' or esc to go back.\n"

	if m.lastAction != "" {
		s += "\n" + m.lastAction + "\n"
	}

	return s
}
