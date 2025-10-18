package main

import (
	"log"
	"os"

	"github.com/deltron-fr/rss-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}


func main() {
	dataConfig := config.Read()
	programState := state{cfg: &dataConfig}

	m := make(map[string]func(*state, command) error)
	cliCommands := commands{registeredCommands: m}

	cliCommands.register("login", handlerLogin)
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	args := os.Args[1:]

	newCommand := command{name: args[0], args: args[1:]}
	
	err := cliCommands.run(&programState, newCommand)
	if err != nil {
		log.Fatal(err)
	}
}
