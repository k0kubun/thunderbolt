package main

import (
	"fmt"
	"github.com/k0kubun/twitter"
	"log"
)

func updateStatus(account *Account, text string) error {
	return account.Client().UpdateStatus(text)
}

func replyStatus(account *Account, text string, tweetId int64) error {
	return account.Client().ReplyStatus(text, tweetId)
}

func homeTimeline(account *Account) {
	client := account.Client()
	tweets, err := client.HomeTimeline()
	if err != nil {
		log.Fatal(err)
	}

	for _, tweet := range reversedTweets(tweets) {
		fmt.Println(formattedTweet(&tweet))
	}
}

func mentionsTimeline(account *Account) {
	client := account.Client()
	tweets, err := client.MentionsTimeline()
	if err != nil {
		log.Fatal(err)
	}

	for _, tweet := range reversedTweets(tweets) {
		fmt.Println(formattedTweet(&tweet))
	}
}

func userTimeline(account *Account, argument string) {
	client := account.Client()
	tweets, err := client.UserTimeline(argument)
	if err != nil {
		fmt.Printf("'%s' is invalid screen_name\n", argument)
	}

	for _, tweet := range reversedTweets(tweets) {
		fmt.Println(formattedTweet(&tweet))
	}
}

func favorite(account *Account, tweet *twitter.Tweet) error {
	return account.Client().Favorite(tweet.Id)
}

func retweet(account *Account, tweet *twitter.Tweet) error {
	return account.Client().Retweet(tweet.Id)
}

func formattedTweet(tweet *twitter.Tweet) string {
	address := tweetMap.registerTweet(tweet)

	return fmt.Sprintf(
		"%s %s: %s",
		foreGrayText(fmt.Sprintf("[$%s]", address)),
		coloredScreenName(tweet.User.ScreenName),
		tweet.Text,
	)
}

func reversedTweets(tweets []twitter.Tweet) []twitter.Tweet {
	reversed := make([]twitter.Tweet, len(tweets))
	lastIndex := len(tweets) - 1

	for index, tweet := range tweets {
		reversed[lastIndex-index] = tweet
	}
	return reversed
}
