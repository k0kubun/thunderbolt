package main

import (
	"fmt"
	"log"

	"github.com/jessevdk/go-flags"
	"github.com/k0kubun/go-readline"
)

type Options struct {
	ScreenName string `short:"a" long:"account" description:"login as an account of selected screen_name"`
	Tweet      string `short:"t" long:"tweet" description:"just post a tweet and finish"`
}

func main() {
	options := new(Options)
	if _, err := flags.Parse(options); err != nil {
		log.Fatal(err)
	}

	account := loadAccount(options)
	if len(options.Tweet) > 0 {
		updateStatus(account, options.Tweet)
		return
	}

	invokeInteractiveShell(account)
}

func loadAccount(options *Options) *Account {
	if len(options.ScreenName) > 0 {
		return AccountByScreenName(options.ScreenName)
	} else {
		return DefaultAccount()
	}
}

func invokeInteractiveShell(account *Account) {
	readline.CatchSignals(0)

	for {
		currentLine := readline.Readline(prompt(account))
		if currentLine == nil || *currentLine == ":exit" {
			return
		}

		err := executeCommand(account, *currentLine)
		if err != nil {
			fmt.Print(err.Error())
		}
		readline.AddHistory(*currentLine)
	}
}

func prompt(account *Account) *string {
	prompt := fmt.Sprintf("[%s] ", coloredScreenName(account.ScreenName))
	return &prompt
}
