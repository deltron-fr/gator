package main

import (
	"context"
	"fmt"
)

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
