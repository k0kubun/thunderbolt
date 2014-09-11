package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/k0kubun/go-readline"
)

func executeCommand(account *Account, line string) {
	streamBlocked = true
	defer func() { streamBlocked = false }()

	if regexpMatch(line, "^\\$[a-x][a-x] .") {
		confirmReply(account, line[1:3], line[3:])
		return
	} else if !strings.HasPrefix(line, ":") {
		confirmTweet(account, line)
		return
	}

	command, argument := splitCommand(line)
	switch command {
	case "recent":
		recent(account, argument)
	case "mentions":
		mentionsTimeline(account)
	case "favorite":
		confirmFavorite(account, argument)
	case "retweet":
		confirmRetweet(account, argument)
	case "delete":
		confirmDelete(account, argument)
	default:
		commandNotFound()
	}
}

func recent(account *Account, argument string) {
	if len(argument) > 0 {
		userTimeline(account, argument)
	} else {
		homeTimeline(account)
	}
}

func confirmReply(account *Account, address, text string) {
	tweet := tweetMap.tweetByAddress(address)
	if tweet == nil || tweet.Id == 0 {
		println("Tweet is not registered")
		return
	}

	replyText := fmt.Sprintf("@%s%s", tweet.User.ScreenName, text)

	replyTarget := fmt.Sprintf("'%s: %s'", tweet.User.ScreenName, tweet.Text)
	confirmMessage := fmt.Sprintf("reply '%s'", replyText)
	confirmExecute(func() error {
		return replyStatus(account, replyText, tweet.Id)
	}, "%s\n%s", foreGrayText(replyTarget), foreColoredText(confirmMessage, "red"))
}

func confirmTweet(account *Account, text string) {
	confirmExecute(func() error {
		return updateStatus(account, text)
	}, foreColoredText("update '%s'", "red"), text)
}

func confirmFavorite(account *Account, argument string) {
	address := extractAddress(argument)
	if address == "" {
		commandNotFound()
		return
	}

	tweet := tweetMap.tweetByAddress(address)
	if tweet == nil || tweet.Id == 0 {
		println("Tweet is not registered")
		return
	}

	confirmExecute(func() error {
		return favorite(account, tweet)
	}, foreColoredText("favorite '%s'", "red"), tweet.Text)
}

func confirmRetweet(account *Account, argument string) {
	address := extractAddress(argument)
	if address == "" {
		commandNotFound()
		return
	}

	tweet := tweetMap.tweetByAddress(address)
	if tweet == nil || tweet.Id == 0 {
		println("Tweet is not registered")
		return
	}

	confirmExecute(func() error {
		return retweet(account, tweet)
	}, foreColoredText("retweet '%s'", "red"), tweet.Text)
}

func confirmDelete(account *Account, argument string) {
	address := extractAddress(argument)
	if address == "" {
		commandNotFound()
		return
	}

	tweet := tweetMap.tweetByAddress(address)
	if tweet == nil || tweet.Id == 0 {
		println("Tweet is not registered")
		return
	}

	confirmExecute(func() error {
		return delete(account, tweet)
	}, foreColoredText("delete '%s'", "red"), tweet.Text)
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

func commandNotFound() {
	fmt.Printf("%s\n", backColoredText(foreBlackText("Command not found"), "yellow"))
}
