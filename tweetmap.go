package main

import (
	"fmt"
	"github.com/k0kubun/twitter"
)

var (
	tweetMap        = NewTweetMapper()
	alphabetNumber  = 26
	maxIndex        = alphabetNumber*alphabetNumber - 1
	inexistentIndex = -1
)

// Singleton struct
type TweetMapper struct {
	tweets    []twitter.Tweet
	lastIndex int // index for latest tweet
}

func NewTweetMapper() *TweetMapper {
	return &TweetMapper{
		tweets:    make([]twitter.Tweet, maxIndex+1),
		lastIndex: maxIndex,
	}
}

func (t *TweetMapper) registerTweet(tweet *twitter.Tweet) string {
	registeredIndex := t.registeredIndex(tweet)
	if registeredIndex != inexistentIndex {
		return t.addressByIndex(registeredIndex)
	}

	newIndex := t.incrementIndex(t.lastIndex)
	t.tweets[newIndex] = *tweet
	t.lastIndex = newIndex
	return t.addressByIndex(newIndex)
}

func (t *TweetMapper) incrementIndex(index int) int {
	if index >= maxIndex {
		return 0
	} else {
		return index + 1
	}
}

func (t *TweetMapper) registeredIndex(tweet *twitter.Tweet) int {
	for index, registeredTweet := range t.tweets {
		if registeredTweet.Id == tweet.Id {
			return index
		}
	}
	return inexistentIndex
}

func (t *TweetMapper) indexByAddress(address string) int {
	return 0
}

func (t *TweetMapper) tweetById(id int64) *twitter.Tweet {
	for _, tweet := range t.tweets {
		if tweet.Id == id {
			return &tweet
		}
	}
	return nil
}

func (t *TweetMapper) addressByIndex(index int) string {
	lowerIndex := index % alphabetNumber
	higherIndex := index / alphabetNumber
	return fmt.Sprintf("%c%c", 'a'+higherIndex, 'a'+lowerIndex)
}
