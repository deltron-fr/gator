package main

import (
	"context"
	"fmt"
	"time"

	"github.com/deltron-fr/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command, user database.User) error {
	if len(cmd.args) < 2 {
		return fmt.Errorf("usage: %s <feed_name> <feed_url>", cmd.name)
	}

	feedName, feedURL := cmd.args[0], cmd.args[1]
	
	feed, err := s.db.Createfeed(context.Background(), database.CreatefeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		Name:      feedName,
		Url:       feedURL,
	})

	if err != nil {
		return fmt.Errorf("couldn't create feed: %w", err)
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    feed.ID,
	})

	if err != nil {
		return fmt.Errorf("could not create feedfollow, err: %w", err)
	}


	fmt.Println("Feed created successfully:")
	printFeed(database.Feed(feed))
	fmt.Println()
	fmt.Println("=====================================")
	return nil
}

func handlerGetFeeds(s *state, cmd command) error {
	if len(cmd.args) != 0 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return fmt.Errorf("err: %w", err)
	}

	for _, feed := range feeds {
		user, err := s.db.GetUserFeeds(context.Background(), feed.UserID)
		if err != nil {
			return fmt.Errorf("unable to get user info, err: %w", err)
		}

		fmt.Printf("Name of feed: %s\n", feed.Name)
		fmt.Printf("URL of feed: %s\n", feed.Url)
		fmt.Printf("User that created feed: %s\n", user)
		fmt.Print("------------------------------\n")
		fmt.Print("\n")
	}

	return nil
}


func printFeed(feed database.Feed) {
	fmt.Printf("* ID:            %s\n", feed.ID)
	fmt.Printf("* Created:       %v\n", feed.CreatedAt)
	fmt.Printf("* Updated:       %v\n", feed.UpdatedAt)
	fmt.Printf("* Name:          %s\n", feed.Name)
	fmt.Printf("* URL:           %s\n", feed.Url)
	fmt.Printf("* UserID:        %s\n", feed.UserID)
}
