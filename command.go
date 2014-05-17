package main

import (
	"fmt"
	"github.com/k0kubun/go-readline"
)

func executeCommand(account *Account, line string) {
	streamBlocked = true

	tweet(account, line)

	streamBlocked = false
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
