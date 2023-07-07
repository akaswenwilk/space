package program

import (
	"fmt"
	"os"

	"github.com/akaswenwilk/space/pkg/configuration"
	"github.com/akaswenwilk/space/pkg/program/newspace"
	"github.com/atotto/clipboard"
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
		os.Exit(1)
	}

	clipboard.WriteAll(model.FinalSpace)
	fmt.Printf("space created and stored to clipboard: %s\n", model.FinalSpace)
}

func Purge(conf configuration.Conf) {
	err := os.RemoveAll(conf.SpacesDirectory)
	if err != nil {
		fmt.Printf("error purging: %v", err)
		os.Exit(1)
	}
	fmt.Println("spaces purged successfully")
}
