package main

import (
	"fmt"
	"github.com/k0kubun/twitter"
	"log"
	"regexp"
	"time"
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
		fmt.Println(timelineSeparator() + formattedTweet(&tweet))
	}
}

func mentionsTimeline(account *Account) {
	client := account.Client()
	tweets, err := client.MentionsTimeline()
	if err != nil {
		log.Fatal(err)
	}

	for _, tweet := range reversedTweets(tweets) {
		fmt.Println(timelineSeparator() + formattedTweet(&tweet))
	}
}

func userTimeline(account *Account, argument string) {
	client := account.Client()
	tweets, err := client.UserTimeline(argument)
	if err != nil {
		fmt.Printf("'%s' is invalid screen_name\n", argument)
	}

	for _, tweet := range reversedTweets(tweets) {
		fmt.Println(timelineSeparator() + formattedTweet(&tweet))
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
		"%s %s: %s%s%s",
		foreGrayText(fmt.Sprintf("[$%s]", address)),
		coloredScreenName(tweet.User.ScreenName),
		highlightedTweet(tweet.Text),
		protectedBadge(tweet.User),
		foreGrayText(
			formattedTime(tweet.CreatedAt),
			" - ",
			trimTag(tweet.Source),
		),
	)
}

func highlightedTweet(text string) string {
	re, _ := regexp.Compile("@[a-zA-Z0-9_]+")
	text = re.ReplaceAllStringFunc(text, func(word string) string { return coloredScreenName(word) })

	re, _ = regexp.Compile("https?://[a-zA-Z0-9-_./]+")
	text = re.ReplaceAllStringFunc(text, func(word string) string {
		return underline(foreColoredText(word, "cyan"))
	})
	return text
}

func protectedBadge(user *twitter.User) string {
	if user.Protected {
		return foreColoredText(" [P] ", "red")
	} else {
		return " "
	}
}

func formattedTime(timeText string) string {
	t, err := time.Parse(time.RubyDate, timeText)
	if err != nil {
		log.Fatal(err)
	}
	localTime := t.Local()

	return fmt.Sprintf(
		"%d %s %02d:%02d",
		localTime.Day(),
		localTime.Month(),
		localTime.Hour(),
		localTime.Minute(),
	)
}

func trimTag(text string) string {
	re, err := regexp.Compile("<.+?>")
	if err != nil {
		log.Fatal(err)
	}
	return re.ReplaceAllString(text, "")
}

func reversedTweets(tweets []twitter.Tweet) []twitter.Tweet {
	reversed := make([]twitter.Tweet, len(tweets))
	lastIndex := len(tweets) - 1

	for index, tweet := range tweets {
		reversed[lastIndex-index] = tweet
	}
	return reversed
}

func timelineSeparator() string {
	seconds := time.Now().Second()
	return randomBackColoredText(" ", int(seconds))
}
