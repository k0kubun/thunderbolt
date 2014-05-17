package main

import (
	"github.com/k0kubun/go-readline"
	"github.com/wsxiaoys/terminal/color"
)

func executeCommand(account *Account, line string) {
	streamBlocked = true

	tweet(account, line)

	streamBlocked = false
}

func tweet(account *Account, text string) {
	for {
		color.Printf("@rupdate '%s'\n", text)

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
