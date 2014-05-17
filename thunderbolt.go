package main

import (
	"github.com/fiorix/go-readline"
)

var (
	prompt = "⚡ "
)

func main() {
	for {
		currentLine := readline.Readline(&prompt)
		if currentLine == nil || *currentLine == ":exit" {
			return
		}

		readline.AddHistory(*currentLine)
	}
}
