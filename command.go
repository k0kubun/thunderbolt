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
		tweet(account, line)
		return
	}

	command, _ := splitCommand(line)
	switch command {
	case "recent":
		recent(account)
	default:
		fmt.Printf("%s\n", backColoredText(foreBlackText("Command not found"), "yellow"))
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

func tweet(account *Account, text string) {
	for {
		notice := fmt.Sprintf("update '%s'\n", text)
		fmt.Printf(foreColoredText(notice, "red"))

		answer := confirm("[Yn] ")
		if answer == "Y" || answer == "y" || answer == "" {
			err := account.Client().UpdateStatus(text)
			if err != nil {
				print(err)
			}
			return
		} else if answer == "N" || answer == "n" {
			return
		}
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
