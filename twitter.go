package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/k0kubun/twitter"
)

var unescape = strings.NewReplacer(
	"&amp;", "&",
	"&lt;", "<",
	"&gt;", ">",
)

func updateStatus(account *Account, text string) error {
	return account.Client().UpdateStatus(text)
}

func replyStatus(account *Account, text string, tweetId int64) error {
	return account.Client().ReplyStatus(text, tweetId)
}

func homeTimeline(account *Account) error {
	client := account.Client()
	tweets, err := client.HomeTimeline()
	if err != nil {
		return err
	}

	for _, tweet := range reversedTweets(tweets) {
		fmt.Println(timelineSeparator() + formattedTweet(&tweet))
	}
	return nil
}

func mentionsTimeline(account *Account) error {
	client := account.Client()
	tweets, err := client.MentionsTimeline()
	if err != nil {
		return err
	}

	for _, tweet := range reversedTweets(tweets) {
		fmt.Println(timelineSeparator() + formattedTweet(&tweet))
	}
	return nil
}

func userTimeline(account *Account, argument string) error {
	client := account.Client()
	tweets, err := client.UserTimeline(argument)
	if err != nil {
		return err
	}

	for _, tweet := range reversedTweets(tweets) {
		fmt.Println(timelineSeparator() + formattedTweet(&tweet))
	}
	return nil
}

func favorite(account *Account, tweet *twitter.Tweet) error {
	return account.Client().Favorite(tweet.Id)
}

func retweet(account *Account, tweet *twitter.Tweet) error {
	return account.Client().Retweet(tweet.Id)
}

func delete(account *Account, tweet *twitter.Tweet) error {
	return account.Client().Destroy(tweet.Id)
}

func getLists(account *Account) error {
	client := account.Client()
	lists, err := client.Lists()
	if err != nil {
		return err
	}
	for _, list := range lists {
		//FullName is "@:screen_name/slug"
		fmt.Println(list.FullName[1:])
	}
	return nil
}

func listTimeline(account *Account, argument string) error {
	fullName := strings.SplitN(argument, "/", 2)
	ownerScreenName := fullName[0]
	slug := fullName[1]
	client := account.Client()
	tweets, err := client.ListTimeline(ownerScreenName, slug)
	if err != nil {
		return err
	}

	for _, tweet := range reversedTweets(tweets) {
		fmt.Println(timelineSeparator() + formattedTweet(&tweet))
	}
	return nil
}

func search(account *Account, query string) error {
	client := account.Client()
	tweets, err := client.Search(query)
	if err != nil {
		return err
	}

	for _, tweet := range reversedTweets(tweets) {
		fmt.Println(timelineSeparator() + formattedTweet(&tweet))
	}

	return nil
}

func formattedTweet(tweet *twitter.Tweet) string {
	address := tweetMap.registerTweet(tweet)
	header := fmt.Sprintf(
		"%s %s:",
		foreGrayText(fmt.Sprintf("[$%s]", address)),
		coloredScreenName(tweet.User.ScreenName),
	)
	footer := fmt.Sprintf(
		"%s%s",
		protectedBadge(tweet.User),
		foreGrayText(
			formattedTime(tweet.CreatedAt),
			" - ",
			trimTag(tweet.Source),
		),
	)

	if regexpMatch(tweet.Text, "\n") {
		return fmt.Sprintf(
			"%s%s\n       %s",
			header,
			liftedTweet(highlightedTweet(tweet.Text)),
			footer,
		)
	} else {
		return fmt.Sprintf(
			"%s %s %s",
			header,
			highlightedTweet(unescape.Replace(tweet.Text)),
			footer,
		)
	}
}

func liftedTweet(text string) string {
	re, _ := regexp.Compile("(^|\n)")
	return re.ReplaceAllString(
		text,
		fmt.Sprintf("\n       %s", foreGrayText("|")),
	)
}

func highlightedTweet(text string) string {
	// Highlight screen_name
	re, _ := regexp.Compile("@[a-zA-Z0-9_]+")
	text = re.ReplaceAllStringFunc(text, func(word string) string { return coloredScreenName(word) })

	// Highlight URL
	re, _ = regexp.Compile("https?://[a-zA-Z0-9-_./]+")
	text = re.ReplaceAllStringFunc(text, func(word string) string {
		return underline(foreColoredText(word, "cyan"))
	})

	// Highlight hash tag
	re, _ = regexp.Compile("#[^ ]+")
	text = re.ReplaceAllStringFunc(text, func(word string) string { return coloredScreenName(word) })

	return text
}

func protectedBadge(user *twitter.User) string {
	if user.Protected {
		return foreColoredText("[P] ", "red")
	} else {
		return ""
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
