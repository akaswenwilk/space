package newspace

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/akaswenwilk/space/pkg/configuration"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/go-git/go-git/v5"
	"github.com/lithammer/fuzzysearch/fuzzy"
)

type Model struct {
	Conf           *configuration.Conf
	Text           []string
	Repo           string
	RepoSelected   bool
	Branch         string
	BranchSelected bool
	Err            error
	FinalSpace     string
	SelectedRepo   int
}

func Start(conf *configuration.Conf) *Model {
	return &Model{Conf: conf}
}

func (m *Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}

func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
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
				m.SelectedRepo = 0
			}

		case messageString == "down" || messageString == "tab":
			m.SelectedRepo += 1
			if m.SelectedRepo > len(m.Repos()) {
				m.SelectedRepo = 1
			}

		case messageString == "up":
			m.SelectedRepo -= 1
			if m.SelectedRepo < 1 {
				m.SelectedRepo = len(m.Repos())
			}

		// submit text
		case messageString == "enter":
			if m.RepoSelected {
				m.Branch = strings.Join(m.Text, "")
				m.BranchSelected = true
			} else {
				if m.SelectedRepo == 0 {
					m.Repo = strings.Join(m.Text, "")
				} else {
					m.Repo = m.Repos()[m.SelectedRepo-1]
				}
				m.RepoSelected = true
			}

			m.Text = []string{}
			if m.RepoSelected && m.BranchSelected {
				space, err := m.Clone()
				if err != nil && !errors.Is(err, git.ErrRepositoryAlreadyExists) {
					m.Err = fmt.Errorf("error cloning: %w", err)
				}
				m.FinalSpace = space
			}

		case len(messageString) == 1:
			m.Text = append(m.Text, messageString)
			m.SelectedRepo = 0
		}
	}

	if m.Err != nil || m.FinalSpace != "" {
		return m, tea.Quit
	}

	// Return the updated model to the Bubble Tea runtime for processing.
	// Note that we're not returning a command.
	return m, nil
}

func (m *Model) View() string {
	// Send the UI for rendering
	switch {
	case !m.RepoSelected:
		return fmt.Sprintf("Repo: %s\n\n%s", strings.Join(m.Text, ""), m.PossibleRepos())
	case m.RepoSelected:
		return fmt.Sprintf("Repo: %s\nBranch (leave blank for default branch): %s", m.Repo, strings.Join(m.Text, ""))
	default:
		return ""
	}
}

func (m *Model) PossibleRepos() string {
	str := ""
	for i, r := range m.Repos() {
		marker := " "
		if i == m.SelectedRepo-1 {
			marker = "x"
		}
		str += fmt.Sprintf("[%s] %s\n", marker, r)
	}
	return str
}

func (m *Model) Repos() []string {
	entry := strings.Join(m.Text, "")
	if entry == "" {
		if len(m.Conf.Spaces) > 10 {
			return m.Conf.Spaces[:10]
		}

		return m.Conf.Spaces
	}

	matches := fuzzy.RankFindFold(entry, m.Conf.Spaces)

	sort.Sort(matches)

	var result []string

	for i, m := range matches {
		if i > 10 {
			break
		}

		result = append(result, m.Target)
	}

	return result
}
