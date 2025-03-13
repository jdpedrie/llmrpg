package pages

import (
	fmt "fmt"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type WizardModel struct {
	step      int
	nameInput textinput.Model
	ageInput  textinput.Model
}

func NewWizardModel() *WizardModel {
	// Set up the text inputs
	nameInput := textinput.New()
	nameInput.Placeholder = "Type your name"
	nameInput.Focus()

	ageInput := textinput.New()
	ageInput.Placeholder = "Type your age"

	return &WizardModel{
		step:      0,
		nameInput: nameInput,
		ageInput:  ageInput,
	}
}

func (m *WizardModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m *WizardModel) Update(msg tea.Msg) (Page, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch m.step {
		case 0:
			// Step 1: Enter name
			var cmd tea.Cmd
			m.nameInput, cmd = m.nameInput.Update(msg)
			if msg.String() == "enter" && len(m.nameInput.Value()) > 0 {
				// Move to step 2
				m.step = 1
				m.nameInput.Blur()
				m.ageInput.Focus()
			}
			return m, cmd

		case 1:
			// Step 2: Enter age
			var cmd tea.Cmd
			m.ageInput, cmd = m.ageInput.Update(msg)
			if msg.String() == "enter" && len(m.ageInput.Value()) > 0 {
				// Move to step 3
				m.step = 2
				m.ageInput.Blur()
			}
			return m, cmd

		case 2:
			// Step 3: Summary
			// Pressing Enter, q, or esc returns to main menu
			switch msg.String() {
			case "enter", "q", "esc":
				// Reset wizard if you want, or keep the data
				return m, func() tea.Msg {
					return SwitchPageMsg{PageName: "mainmenu"}
				}
			}
		}
	}
	return m, nil
}

func (m *WizardModel) View() string {
	switch m.step {
	case 0:
		return fmt.Sprintf(
			"Wizard Step 1: Enter your name\n\n%s\n\n(press Enter to continue)",
			m.nameInput.View(),
		)
	case 1:
		return fmt.Sprintf(
			"Wizard Step 2: Enter your age\n\n%s\n\n(press Enter to continue)",
			m.ageInput.View(),
		)
	case 2:
		return fmt.Sprintf(
			"Name: %s\nAge: %s\n\nThanks for the info!\n\n(press Enter, 'q', or ESC to return to main menu)",
			m.nameInput.Value(), m.ageInput.Value(),
		)
	}
	return ""
}
