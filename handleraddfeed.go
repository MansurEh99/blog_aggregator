package main

import (
	"fmt"
	"context"
	"time"
	"errors"

	"github.com/google/uuid"
	"github.com/MansurEh99/blog_aggregator/internal/database"
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

	fmt.Println(feed.ID)
	fmt.Println(feed.CreatedAt)
	fmt.Println(feed.UpdatedAt)
	fmt.Println(feed.Name)
	fmt.Println(feed.Url)
	fmt.Println(feed.UserID)

	return nil
}
