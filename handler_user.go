package main

import (
	"fmt"

)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	userName := cmd.args[0]
	err := s.cfg.SetUser(userName)
	if err != nil {
		return fmt.Errorf("couldn't set username")
	}

	fmt.Printf("%s has been set as user", userName)
	return nil
}