package main

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/k0kubun/go-readline"
)

var (
	commandNotFound = errors.New(
		fmt.Sprintf("%s\n", backColoredText(foreBlackText("Command not found"), "yellow")),
	)
)

func executeCommand(account *Account, line string) error {
	streamBlocked = true
	defer func() { streamBlocked = false }()

	if regexpMatch(line, "^\\$[a-x][a-x] .") {
		return confirmReply(account, line[1:3], line[3:])
	} else if !strings.HasPrefix(line, ":") {
		confirmTweet(account, line)
		return nil
	}

	command, argument := splitCommand(line)
	switch command {
	case "recent":
		return recent(account, argument)
	case "mentions":
		return mentionsTimeline(account)
	case "favorite":
		return confirmFavorite(account, argument)
	case "retweet":
		return confirmRetweet(account, argument)
	case "delete":
		return confirmDelete(account, argument)
	default:
		return commandNotFound
	}
}

func recent(account *Account, argument string) error {
	if len(argument) > 0 {
		return userTimeline(account, argument)
	} else {
		return homeTimeline(account)
	}
}

func confirmReply(account *Account, address, text string) error {
	tweet, err := tweetMap.tweetByAddress(address)
	if err != nil {
		return err
	}

	replyText := fmt.Sprintf("@%s%s", tweet.User.ScreenName, text)

	replyTarget := fmt.Sprintf("'%s: %s'", tweet.User.ScreenName, tweet.Text)
	confirmMessage := fmt.Sprintf("reply '%s'", replyText)
	confirmExecute(func() error {
		return replyStatus(account, replyText, tweet.Id)
	}, "%s\n%s", foreGrayText(replyTarget), foreColoredText(confirmMessage, "red"))

	return nil
}

func confirmTweet(account *Account, text string) {
	confirmExecute(func() error {
		return updateStatus(account, text)
	}, foreColoredText("update '%s'", "red"), text)
}

func confirmFavorite(account *Account, argument string) error {
	address := extractAddress(argument)
	if address == "" {
		return commandNotFound
	}

	tweet, err := tweetMap.tweetByAddress(address)
	if err != nil {
		return err
	}

	confirmExecute(func() error {
		return favorite(account, tweet)
	}, foreColoredText("favorite '%s'", "red"), tweet.Text)

	return nil
}

func confirmRetweet(account *Account, argument string) error {
	address := extractAddress(argument)
	if address == "" {
		return commandNotFound
	}

	tweet, err := tweetMap.tweetByAddress(address)
	if err != nil {
		return err
	}

	confirmExecute(func() error {
		return retweet(account, tweet)
	}, foreColoredText("retweet '%s'", "red"), tweet.Text)

	return nil
}

func confirmDelete(account *Account, argument string) error {
	address := extractAddress(argument)
	if address == "" {
		return commandNotFound
	}

	tweet, err := tweetMap.tweetByAddress(address)
	if err != nil {
		return err
	}

	confirmExecute(func() error {
		return delete(account, tweet)
	}, foreColoredText("delete '%s'", "red"), tweet.Text)

	return nil
}

func confirmExecute(function func() error, format string, a ...interface{}) {
	confirmMessage := fmt.Sprintf(format, a...)

	for {
		fmt.Println(confirmMessage)

		answer := excuse("[Yn] ")
		if answer == "Y" || answer == "y" || answer == "" {
			err := function()
			if err != nil {
				fmt.Printf("Error: %s\n", err.Error())
			}
			return
		} else if answer == "N" || answer == "n" {
			return
		}
	}
}

func excuse(prompt string) string {
	result := readline.Readline(&prompt)
	if result == nil {
		print("\n")
		return "n"
	}
	return *result
}

func regexpMatch(text string, exp string) bool {
	re, err := regexp.Compile(exp)
	if err != nil {
		log.Fatal(err)
	}
	return re.MatchString(text)
}

func splitCommand(text string) (string, string) {
	re, err := regexp.Compile("^:[^ ]+")
	if err != nil {
		log.Fatal(err)
	}

	result := re.FindStringIndex(text)
	if result == nil {
		return text[1:], ""
	}
	last := result[1]

	if last+1 >= len(text) {
		return text[1:], ""
	}
	return text[1:last], text[last+1:]
}

func extractAddress(argument string) string {
	re, err := regexp.Compile("\\$[a-z][a-z]")
	if err != nil {
		log.Fatal(err)
	}

	result := re.FindString(argument)
	if result == "" {
		return ""
	} else {
		return result[1:]
	}
}
