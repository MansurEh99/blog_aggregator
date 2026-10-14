package main

import (
	"errors"
)

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.commands[cmd.name]
	if !ok {
		return errors.New("invalid command")
	}

	return f(s, cmd)
}
