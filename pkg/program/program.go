package program

import (
	"fmt"
	"os"

	"github.com/akaswenwilk/space/pkg/configuration"
	"github.com/akaswenwilk/space/pkg/program/newspace"
	tea "github.com/charmbracelet/bubbletea"
)

func New(conf configuration.Conf) {
	p := tea.NewProgram(newspace.Start(conf))
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func Purge(conf configuration.Conf) {
}
