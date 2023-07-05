package model

import (
	"fmt"
	"regexp"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

var charRegex = regexp.MustCompile(`[a-zA-Z0-9-_]{1}`)

type Model struct {
	Text []string
}

func New() Model {
	return Model{}
}

func (m Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// Is it a key press?
	case tea.KeyMsg:
		messageString := msg.String()

		switch {

		// These keys should exit the program.
		case messageString == "ctrl+c":
			return m, tea.Quit

		// handle clear text
		case messageString == "ctrl+u":
			m.Text = []string{}

		// remove end character
		case messageString == "backspace":
			if len(m.Text) > 0 {
				m.Text = m.Text[:len(m.Text)-1]
			}

		// submit text
		case messageString == "enter":
			// TODO: handle enter

		case len(messageString) == 1:
			m.Text = append(m.Text, messageString)
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m Model) View() string {
	// Send the UI for rendering
	return fmt.Sprintf("Repo Name: %s", strings.Join(m.Text, ""))
}
