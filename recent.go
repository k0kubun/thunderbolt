package main

import (
	"fmt"
	"github.com/k0kubun/twitter"
	"log"
)

func recent(account *Account, argument string) {
	if len(argument) > 0 {
		userTimeline(account, argument)
	} else {
		homeTimeline(account)
	}
}

func homeTimeline(account *Account) {
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

func userTimeline(account *Account, argument string) {
	client := account.Client()
	tweets, err := client.UserTimeline(argument)
	if err != nil {
		fmt.Printf("'%s' is invalid screen_name\n", argument)
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
