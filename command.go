package main

import (
	"fmt"
	"github.com/k0kubun/go-readline"
	"log"
	"regexp"
	"strings"
)

func executeCommand(account *Account, line string) {
	streamBlocked = true
	defer func() { streamBlocked = false }()

	if !strings.HasPrefix(line, ":") {
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
	default:
		commandNotFound()
	}
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

func confirmTweet(account *Account, text string) {
	for {
		notice := fmt.Sprintf("update '%s'\n", text)
		fmt.Printf(foreColoredText(notice, "red"))

		answer := confirm("[Yn] ")
		if answer == "Y" || answer == "y" || answer == "" {
			err := updateStatus(account, text)
			if err != nil {
				print(err)
			}
			return
		} else if answer == "N" || answer == "n" {
			return
		}
	}
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

	for {
		notice := fmt.Sprintf("favorite '%s'\n", tweet.Text)
		fmt.Printf(foreColoredText(notice, "red"))

		answer := confirm("[Yn] ")
		if answer == "Y" || answer == "y" || answer == "" {
			err := favorite(account, tweet)
			if err != nil {
				print(err)
			}
			return
		} else if answer == "N" || answer == "n" {
			return
		}
	}
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

func confirm(prompt string) string {
	result := readline.Readline(&prompt)
	if result == nil {
		print("\n")
		return "n"
	}
	return *result
}

func commandNotFound() {
	fmt.Printf("%s\n", backColoredText(foreBlackText("Command not found"), "yellow"))
}
