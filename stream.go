package main

import (
	"fmt"
	"github.com/k0kubun/go-readline"
	"github.com/k0kubun/userstream"
)

var (
	lineQueue = []string{}
)

func startUserStream(account *Account) {
	client := &userstream.Client{
		ConsumerKey:       account.ConsumerKey,
		ConsumerSecret:    account.ConsumerSecret,
		AccessToken:       account.AccessToken,
		AccessTokenSecret: account.AccessTokenSecret,
	}
	go client.UserStream(printEvent)
}

func printEvent(event interface{}) {
	switch event.(type) {
	case *userstream.Tweet:
		tweet := event.(*userstream.Tweet)
		insertLine("%s: %s", tweet.User.ScreenName, tweet.Text)
	case *userstream.Delete:
		tweetDelete := event.(*userstream.Delete)
		insertLine("[delete] %d", tweetDelete.Id)
	case *userstream.Favorite:
		favorite := event.(*userstream.Favorite)
		insertLine("[favorite] %s => %s : %s",
			favorite.Source.ScreenName, favorite.Target.ScreenName, favorite.TargetObject.Text)
	case *userstream.Unfavorite:
		unfavorite := event.(*userstream.Unfavorite)
		insertLine("[unfavorite] %s => %s : %s",
			unfavorite.Source.ScreenName, unfavorite.Target.ScreenName, unfavorite.TargetObject.Text)
	case *userstream.Follow:
		follow := event.(*userstream.Follow)
		insertLine("[follow] %s => %s", follow.Source.ScreenName, follow.Target.ScreenName)
	case *userstream.Unfollow:
		unfollow := event.(*userstream.Unfollow)
		insertLine("[unfollow] %s => %s", unfollow.Source.ScreenName, unfollow.Target.ScreenName)
	case *userstream.ListMemberAdded:
		listMemberAdded := event.(*userstream.ListMemberAdded)
		insertLine("[list_member_added] %s (%s)",
			listMemberAdded.TargetObject.FullName, listMemberAdded.TargetObject.Description)
	case *userstream.ListMemberRemoved:
		listMemberRemoved := event.(*userstream.ListMemberRemoved)
		insertLine("[list_member_removed] %s (%s)",
			listMemberRemoved.TargetObject.FullName, listMemberRemoved.TargetObject.Description)
	}
}

func insertLine(format string, a ...interface{}) {
	line := fmt.Sprintf(format, a...)
	lineQueue = append(lineQueue, line)

	if len(readline.LineBuffer()) == 0 {
		fmt.Printf("\033[0G\033[K")
		for _, line := range lineQueue {
			fmt.Println(line)
		}
		lineQueue = []string{}
		readline.RefreshLine()
	}
}
