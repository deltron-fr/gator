package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/deltron-fr/gator/internal/database"
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

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	userName := cmd.args[0]
	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return errors.New("user does not exist")
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return fmt.Errorf("couldn't set username")
	}

	fmt.Printf("%s has been set as user", user.Name)
	return nil
}

func handlerUsers(s *state, cmd command) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error getting users, err: %v", err)
	}

	currentUser := s.cfg.CurrentUsername
	for _, user := range users {
		if user.Name == currentUser {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}

func handlerReset(s *state, cmd command) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	err := s.db.ResetDB(context.Background())
	if err != nil {
		return fmt.Errorf("error resetting database: %v", err)
	}

	return nil
}
