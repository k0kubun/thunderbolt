package main

import (
	"fmt"
	"github.com/k0kubun/twitter"
	"log"
)

func recent(account *Account) {
	client := account.Client()
	tweets, err := client.HomeTimeline()
	if err != nil {
		log.Fatal(err)
	}

	for _, tweet := range reversedTweets(tweets) {
		fmt.Printf(
			"%s: %s\n",
			coloredScreenName(tweet.User.ScreenName),
			tweet.Text,
		)
	}
}

func reversedTweets(tweets []twitter.Tweet) []twitter.Tweet {
	reversed := make([]twitter.Tweet, len(tweets))
	lastIndex := len(tweets) - 1

	for index, tweet := range tweets {
		reversed[lastIndex-index] = tweet
	}
	return reversed
}
