package main

import (
	"fmt"
	"context"
	"errors"
	"strconv"

	"github.com/MansurEh99/gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	limit := 2
	if len(cmd.args) > 0 {
		converted, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return errors.New("input a number")
		}
		limit = converted
	}

	arg := database.GetPostsForUserParams {
		ID: user.ID,
		Limit: limit,
	}

	posts, err := s.db.GetPostsForUser(context.Background(), arg)
	if err != nil {
		return errors.New("unable to get posts")
	}

	for _, post := range posts {
		fmt.Println("Post:")
		fmt.Printf("- id: %v\n", post.ID)
		fmt.Printf("- created at: %v\n", post.CreatedAt)
		fmt.Printf("- updated at: %v\n", post.UpdatedAt)
		if post.Title.Valid {
			fmt.Printf("- title: %v\n", post.Title.String)
		}
		fmt.Printf("- url: %v\n", post.Url)
		if post.Description.Valid {
			fmt.Printf("- description: %v\n", post.Description.String)
		}
		if post.PublishedAt.Valid {
			fmt.Printf("- published at: %v\n", post.PublishedAt.Time)
		}
		fmt.Printf("- feed id: %v\n", post.FeedID)
	}

	return nil
}
