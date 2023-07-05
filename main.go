package main

import (
	"log"
	"os"

	"github.com/akaswenwilk/space/pkg/program"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatalf("please enter the program name: e.g. space new")
	}

	switch args[1] {
	case "new":
		program.New()
	case "purge":
		program.Purge()
	default:
		log.Fatalf("available programs are purge and new")
	}
}
