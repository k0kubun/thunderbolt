package main

type Account struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
	ScreenName        string
}

func DefaultAccount() *Account {
	return nil
}

func AccountByScreenName(screenName string) *Account {
	return nil
}
