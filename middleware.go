package main

import (
	"context"
	"fmt"

	"github.com/deltron-fr/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		username := s.cfg.CurrentUsername
		user, err := s.db.GetUser(context.Background(), username)
		if err != nil {
			return fmt.Errorf("could not get user: %w", err)
		}

		err = handler(s, cmd, user)
		if err != nil {
			return err
		}
		return nil
	}
}