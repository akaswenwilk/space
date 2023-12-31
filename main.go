package main

import (
	"errors"
	"log"
	"os"

	"github.com/akaswenwilk/space/pkg/configuration"
	"github.com/akaswenwilk/space/pkg/program"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		log.Fatalf("please enter the program name: e.g. space new")
	}

	conf, err := configuration.New()
	failOnError(err)

	switch args[1] {
	case "new":
		program.New(conf)
	case "purge":
		program.Purge(conf)
	default:
		failOnError(errors.New("available programs are purge and new"))
	}
}

func failOnError(err error) {
	if err != nil {
		log.Fatalf(err.Error())
	}
}
