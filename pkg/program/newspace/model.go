package newspace

import (
	"errors"
	"fmt"
	"strings"

	"github.com/akaswenwilk/space/pkg/configuration"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-git/go-git/v5"
)

type Model struct {
	Conf           configuration.Conf
	Text           []string
	Repo           string
	RepoSelected   bool
	Branch         string
	BranchSelected bool
	Err            error
	Success        bool
	Space          string
}

func Start(conf configuration.Conf) Model {
	return Model{Conf: conf}
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
			if m.RepoSelected {
				m.Branch = strings.Join(m.Text, "")
				m.BranchSelected = true
			} else {
				m.Repo = strings.Join(m.Text, "")
				m.RepoSelected = true
			}

			m.Text = []string{}
			if m.RepoSelected && m.BranchSelected {
				space, err := m.Clone()
				if err != nil && !errors.Is(err, git.ErrRepositoryAlreadyExists) {
					m.Err = fmt.Errorf("error cloning: %w", err)
					return m, nil
				}
				m.Space = space
				m.Success = true
			}

		case len(messageString) == 1:
			m.Text = append(m.Text, messageString)
		}
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m Model) View() string {
	if m.Success {
		return fmt.Sprintf("cd %s", m.Space)
	}
	// Send the UI for rendering
	switch {
	case m.Err != nil:
		return m.Err.Error()
	case m.RepoSelected && m.BranchSelected:
		return fmt.Sprintf("Selected Repo: %s\n\nSelected Branch: %s", m.Repo, m.Branch)
	case m.RepoSelected:
		return fmt.Sprintf("Branch (leave blank for default branch): %s", strings.Join(m.Text, ""))
	default:
		return fmt.Sprintf("Repo: %s", strings.Join(m.Text, ""))
	}
}
