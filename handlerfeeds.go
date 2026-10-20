package main

import (
	"errors"
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return errors.New("unable to get feeds")
	}

	for _, feed := range feeds {
		fmt.Println("feed:")
		fmt.Printf("- name: %v\n", feed.Name)
		fmt.Printf("- url: %v\n", feed.Url)
		fmt.Printf("- username: %v\n", feed.UserName)
	}

	return nil
}
