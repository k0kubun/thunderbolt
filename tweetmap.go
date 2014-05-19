package main

import (
	"github.com/k0kubun/twitter"
)

var (
	tweetMap       = NewTweetMapper()
	alphabetNumber = 26
	maxIndex       = alphabetNumber * alphabetNumber
)

// Singleton struct
type TweetMapper struct {
	tweets    []*twitter.Tweet
	lastIndex int
}

func NewTweetMapper() *TweetMapper {
	return &TweetMapper{}
}
