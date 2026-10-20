package main

import (
	"fmt"
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/MansurEh99/gator/internal/database"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("no url provided")
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return errors.New("cannot find feed")
	}

	arg := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:	user.ID,
		FeedID:	feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), arg)
	if err != nil {
		return errors.New("unable to follow feed")
	}

	fmt.Println("Follow successful!")
	fmt.Printf("feed name: %v\n", feed.Name)
	fmt.Printf("current user: %v\n", user.Name)

	return nil
}
