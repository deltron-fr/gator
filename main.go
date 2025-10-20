package main

import _ "github.com/lib/pq"

import (
	"log"
	"os"
	"database/sql"

	"github.com/deltron-fr/rss-aggregator/internal/config"
	"github.com/deltron-fr/rss-aggregator/internal/database"
)

const dbURL = "postgres://postgres:postgres@localhost:5432/gator"

type state struct {
	db *database.Queries
	cfg *config.Config
}


func main() {
	dataConfig := config.Read()

	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal(err)
	}
	dbQueries := database.New(db)
	programState := state{db: dbQueries, cfg: &dataConfig}

	m := make(map[string]func(*state, command) error)
	cliCommands := commands{registeredCommands: m}

	cliCommands.register("login", handlerLogin)
	cliCommands.register("register", handlerRegister)
	cliCommands.register("reset", handlerReset)
	if len(os.Args) < 2 {
		log.Fatal("Usage: cli <command> [args...]")
	}
	args := os.Args[1:]

	newCommand := command{name: args[0], args: args[1:]}
	
	err = cliCommands.run(&programState, newCommand)
	if err != nil {
		log.Fatal(err)
	}
}
