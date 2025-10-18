package main

import "errors"

type command struct {
	name string
	args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}


func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.name]
	if !ok {
		return errors.New("command does not exist")
	}

	return f(s, cmd)
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}