package main

import (
	"strings"
	"context"
	"errors"
	"time"
	"log"
	"database/sql"

	"github.com/google/uuid"
	"github.com/MansurEh99/gator/internal/database"
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
		validPub := true
		t, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			validPub = false
		}

		validTitle := true
		title := item.Title
		if title == "" {
			validTitle = false
		}

		validDesc := true
		desc := item.Description
		if desc == "" {
			validDesc = false
		}

		arg := database.CreatePostParams {
			ID: uuid.New(),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
			Title:	sql.NullString{
				String: title,
				Valid: validTitle,
			},
			Url:	item.Link,
			Description: sql.NullString{
				String: desc,
				Valid: validDesc,
			},
			PublishedAt: sql.NullTime{
				Time: t,
				Valid: validPub,
			},
			FeedID: nextFeed.ID,
		}

		_, err = s.db.CreatePost(context.Background(), arg)

		if err != nil {
			if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
				continue
			}
			log.Printf("Couldn't create post: %v", err)
			continue
		}
	}
	log.Printf("Feed %s collected, %v posts found", nextFeed.Name, len(feed.Channel.Item))
	return nil
}
