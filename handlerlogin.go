package main

import(
	"fmt"
	"errors"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no command inputed")
	}

	err := s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Problem setting username: %v", err)
	}

	fmt.Println("User set!")
	return nil
}
