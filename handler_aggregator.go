package main

import (
	"context"
	"fmt"
)


func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) > 1 {
		return fmt.Errorf("usage: %s", cmd.name)
	}

	url := "https://www.wagslane.dev/index.xml"

	feed, err := fetchFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error fetching feed: %w", err)
	}

	fmt.Print(feed)

	return nil

}