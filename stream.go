package main

import (
	"fmt"
	"github.com/k0kubun/go-readline"
	"github.com/k0kubun/userstream"
)

var (
	lineQueue     = []string{}
	streamBlocked = false
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
		printTweet(event.(*userstream.Tweet))
	case *userstream.Delete:
		printDelete(event.(*userstream.Delete))
	case *userstream.Favorite:
		printFavorite(event.(*userstream.Favorite))
	case *userstream.Unfavorite:
		printUnfavorite(event.(*userstream.Unfavorite))
	case *userstream.Follow:
		printFollow(event.(*userstream.Follow))
	case *userstream.Unfollow:
		printUnfollow(event.(*userstream.Unfollow))
	case *userstream.ListMemberAdded:
		printListMemberAdded(event.(*userstream.ListMemberAdded))
	case *userstream.ListMemberRemoved:
		printListMemberRemoved(event.(*userstream.ListMemberRemoved))
	}
}

func printTweet(tweet *userstream.Tweet) {
	insertLine(
		"%s: %s",
		coloredScreenName(tweet.User.ScreenName),
		tweet.Text,
	)
}

func printDelete(tweetDelete *userstream.Delete) {
	insertLine(
		"[delete] %d",
		tweetDelete.Id,
	)
}

func printFavorite(favorite *userstream.Favorite) {
	insertLine(
		"%s %s => %s : %s",
		backColoredText("[favorite]", "green"),
		coloredScreenName(favorite.Source.ScreenName),
		coloredScreenName(favorite.Target.ScreenName),
		favorite.TargetObject.Text,
	)
}

func printUnfavorite(unfavorite *userstream.Unfavorite) {
	insertLine(
		"%s %s => %s : %s",
		backColoredText("[unfavorite]", "green"),
		coloredScreenName(unfavorite.Source.ScreenName),
		coloredScreenName(unfavorite.Target.ScreenName),
		unfavorite.TargetObject.Text,
	)
}

func printFollow(follow *userstream.Follow) {
	insertLine(
		"%s %s => %s",
		backColoredText("[follow]", "green"),
		coloredScreenName(follow.Source.ScreenName),
		coloredScreenName(follow.Target.ScreenName),
	)
}

func printUnfollow(unfollow *userstream.Unfollow) {
	insertLine(
		"%s %s => %s",
		backColoredText("[unfollow]", "green"),
		coloredScreenName(unfollow.Source.ScreenName),
		coloredScreenName(unfollow.Target.ScreenName),
	)
}

func printListMemberAdded(listMemberAdded *userstream.ListMemberAdded) {
	insertLine(
		"%s %s (%s)",
		backColoredText("[list_member_added]", "green"),
		listMemberAdded.TargetObject.FullName,
		listMemberAdded.TargetObject.Description,
	)
}

func printListMemberRemoved(listMemberRemoved *userstream.ListMemberRemoved) {
	insertLine(
		"%s %s (%s)",
		backColoredText("[list_member_removed]", "green"),
		listMemberRemoved.TargetObject.FullName,
		listMemberRemoved.TargetObject.Description,
	)
}

func insertLine(format string, a ...interface{}) {
	line := fmt.Sprintf(format, a...)
	lineQueue = append(lineQueue, line)

	if len(readline.LineBuffer()) == 0 && !streamBlocked {
		fmt.Printf("\033[0G\033[K")
		for _, line := range lineQueue {
			fmt.Println(line)
		}
		lineQueue = []string{}
		readline.RefreshLine()
	}
}
