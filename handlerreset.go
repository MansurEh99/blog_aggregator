package main

import (
	"errors"
	"context"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return errors.New("Unable to reset")
	}
	return err
}
