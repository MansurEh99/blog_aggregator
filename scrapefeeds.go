package main

import (
	"fmt"
	"context"
	"errors"
)

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return errors.New("unable to fetch next feed")
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return errors.New("unable to mark as fetched")
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return errors.New("unable to fetch feed")
	}

	for _, item := range feed.Channel.Item {
		fmt.Printf("- %v\n", item.Title)
	}

	return nil
}
