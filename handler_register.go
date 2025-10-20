package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/deltron-fr/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	name := cmd.args[0]
	user, _ := s.db.GetUser(context.Background(), name)
	emptyUser := database.User{}

	if user != emptyUser {
		return fmt.Errorf("user: %s already exists", name)
	}

	newUser := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      name,
	}

	s.db.CreateUser(context.Background(), newUser)
	err := s.cfg.SetUser(name)
	if err != nil {
		return errors.New("error setting user information")
	}

	log.Println("User created")
	log.Printf("user uuid: %v, username: %s,\ntime of account creation: %v", newUser.ID, newUser.Name, newUser.CreatedAt)

	return nil
}
