package main

import (
	"fmt"
	"context"
	"time"
	"errors"

	"github.com/google/uuid"
	"github.com/MansurEh99/gator/internal/database"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return errors.New("command requires two arguments")
	}

	arg := database.CreateFeedParams{
		ID: uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:	cmd.args[0],
		Url:	cmd.args[1],
		UserID:	user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), arg)
	if err != nil {
		return errors.New("unable to create feed")
	}

	cmd.args = cmd.args[1:]
	err = handlerFollow(s, cmd, user)
	if err != nil {
		return err
	}

	fmt.Printf("- id: %v\n", feed.ID)
	fmt.Printf("- created at: %v\n", feed.CreatedAt)
	fmt.Printf("- updated at: %v\n", feed.UpdatedAt)
	fmt.Printf("- name: %v\n", feed.Name)
	fmt.Printf("url: %v\n", feed.Url)
	fmt.Printf("user id: %v\n", feed.UserID)

	return nil
}
