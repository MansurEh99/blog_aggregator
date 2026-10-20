package main

import (
	"fmt"
	"errors"
	"context"

	"github.com/MansurEh99/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	feeds, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		errors.New("Unable to get followed feeds")
	}

	fmt.Println("You are following:")
	for _, feed := range feeds {
		fmt.Printf("- %v\n", feed.FeedName)
	}
	return nil
}
