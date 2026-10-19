package main

import (
	"fmt"
	"errors"
	"time"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.args) == 0 {
		return errors.New("need to input timer like 1s, 1m, 1h...")
	}

	timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return errors.New("no timer inputed or invalid timer")
	}
	fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)

	ticker := time.NewTicker(timeBetweenRequests)
	for ; ; <-ticker.C {
		fmt.Println("fetching...")
		scrapeFeeds(s)
	}
}
