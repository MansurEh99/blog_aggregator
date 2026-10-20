package main

import (
	"fmt"
	"errors"
	"context"

	"github.com/MansurEh99/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) == 0 {
		return errors.New("no url provided")
	}

	feed, err := s.db.GetFeedByUrl(context.Background(), cmd.args[0])
	if err != nil {
		return errors.New("cannot find feed")
	}

	arg := database.DeleteFollowParams {
		UserID: user.ID,
		FeedID: feed.ID,
	}

	err = s.db.DeleteFollow(context.Background(), arg)
	if err != nil {
		return errors.New("unable to unfollow")
	}

	fmt.Println("Unfollow successful!")
	return nil
}
