package main

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/deltron-fr/gator/internal/database"
	"github.com/google/uuid"
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

func handlerBrowse(s *state, cmd command) error {

	limit := 2

	if len(cmd.args) == 1 {
		n, err := strconv.Atoi(cmd.args[0])
		if err != nil {
			return fmt.Errorf("invalid limit: must be a number")
		}
		limit = n
	} else if len(cmd.args) > 1 {
		return fmt.Errorf("usage: %s <limit>. (default limit = 2)", cmd.name)
	}

	posts, err := s.db.GetPosts(context.Background(), int32(limit))
	if err != nil {
		return fmt.Errorf("error getting posts, err: %w", err)
	}

	for _, post := range posts {

		
		fmt.Printf("* Title: %s\n", post.Title)
		fmt.Printf("* URL: %s\n", post.Url)
		// fmt.Printf("* Feed: %s\n", feed_name)
		fmt.Println()
		fmt.Println("================================")
		fmt.Println()
	}

	return nil

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
		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID: uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Title: feed.Title,
			Url: feed.Link,
			Description: feed.Description,
			PublishedAt: sql.NullString{String:feed.PubDate},
			FeedID: nextFeed.ID,
		})

		if err != nil {
			if strings.Contains(err.Error(), "url") {
				fmt.Println("Url already exists. Skipping...")
			} else {
				return fmt.Errorf("error creating post, err: %w", err)
			}
		}
	}

	return nil

}