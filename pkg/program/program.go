package program

import (
	"fmt"
	"os"

	"github.com/akaswenwilk/space/pkg/configuration"
	"github.com/akaswenwilk/space/pkg/program/newspace"
	tea "github.com/charmbracelet/bubbletea"
)

func New(conf configuration.Conf) {
	model := newspace.Start(&conf)
	p := tea.NewProgram(model)
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

	if model.Err != nil {
		fmt.Printf("Alas, there's been an error: %v", model.Err)
	}

	os.Stdout.WriteString(fmt.Sprintf("cd %s", model.FinalSpace))
}

func Purge(conf configuration.Conf) {
}
