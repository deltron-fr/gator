package main

import (
	"context"
	"fmt"
	"time"

	"github.com/deltron-fr/rss-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}

	url := cmd.args[0]

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

func handlerFollowing(s *state, cmd command, user database.User) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	feedFollows, err := s.db.GetFeedFollows(context.Background(), user.ID)
	if err != nil {
		return fmt.Errorf("err: unable to get user following feeds")
	}

	fmt.Printf("Feeds for user: %s\n", user.Name)
	for i, feedfollow := range feedFollows {
		fmt.Printf("* Feed %d:	%s\n", i+1, feedfollow.FeedName)
	}

	return nil

}