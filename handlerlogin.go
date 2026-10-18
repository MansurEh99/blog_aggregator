package main

import(
	"fmt"
	"errors"
	"os"
	"context"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no command inputed")
	}

	_, err := s.db.GetUser(context.Background(), cmd.args[0])
	if err != nil {
		os.Exit(1)
		return err
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return fmt.Errorf("Problem setting username: %v", err)
	}

	fmt.Println("User set!")
	return nil
}
