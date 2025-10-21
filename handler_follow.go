package main

import (
	"context"
	"fmt"
	"time"

	"github.com/deltron-fr/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	url := cmd.args[0]

	username := s.cfg.CurrentUsername
	user, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("could not get user: %w", err)
	}

	feed, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("could not get feed: %w", err)
	}

	feedFollow, err := s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("could not create feed follow record: %w", err)
	}

	fmt.Printf("current user: %s, name of feed: %s\n", feedFollow.UserName, feedFollow.FeedName)
	return nil
}

func handlerFollowing(s *state, cmd command) error {
	
}