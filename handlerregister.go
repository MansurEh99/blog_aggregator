package main

import (
	"fmt"
	"errors"
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/MansurEh99/blog_aggregator/internal/database"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("no name provided")
	}

	arg := database.CreateUserParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:	cmd.args[0],
	}

	user, err := s.db.CreateUser(context.Background(), arg)
	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
    	return errors.New("name already exists")
	}
	if err != nil {
		return errors.New("unable to create user")
	}

	err = s.cfg.SetUser(cmd.args[0])
	if err != nil {
		return err
	}

	fmt.Println("user created")
	fmt.Printf("- id: %v\n", user.ID)
	fmt.Printf("- created at: %v\n", user.CreatedAt)
	fmt.Printf("- updated at: %v\n", user.UpdatedAt)
	fmt.Printf("- name: %v\n", user.Name)
	return nil
}
