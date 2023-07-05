package program

import (
	"fmt"
	"os"

	"github.com/akaswenwilk/space/pkg/model"
	tea "github.com/charmbracelet/bubbletea"
)

func New() {
	p := tea.NewProgram(model.New())
	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}

func Purge() {
}
