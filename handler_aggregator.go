package main

import (
	"context"
	"fmt"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time eg: 1s, 1m>", cmd.name)
	}

	time_between_reqs := cmd.args[0]
	timeBetweenReqs, err := time.ParseDuration(time_between_reqs)
	if err != nil {
		return fmt.Errorf("error parsing duration, err: %w", err)
	}

	ticker := time.NewTicker(timeBetweenReqs)
	for ; ; <-ticker.C {
		fmt.Printf("Collecting feeds every %v\n", timeBetweenReqs)
		scrapeFeeds(s)
	}
}


func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("error getting next feed, err: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), nextFeed.ID)
	if err != nil {
		return fmt.Errorf("error marking feed, err: %w", err)
	}

	rssfeeds, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("error fetching feed, err: %w", err)
	}

	for _, feed := range rssfeeds.Channel.Item {
		fmt.Printf("- %s\n", feed.Title)
	}

	return nil

}