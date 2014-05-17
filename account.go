package main

import (
	"bytes"
	"encoding/json"
	"github.com/k0kubun/twitter"
	"github.com/k0kubun/twitter-auth/auth"
	"io/ioutil"
	"log"
	"os"
	"os/user"
)

type Account struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
	ScreenName        string

	clnt *twitter.Client
}

func (a *Account) Client() *twitter.Client {
	if a.clnt != nil {
		return a.clnt
	}

	a.clnt = &twitter.Client{
		ConsumerKey:       a.ConsumerKey,
		ConsumerSecret:    a.ConsumerSecret,
		AccessToken:       a.AccessToken,
		AccessTokenSecret: a.AccessTokenSecret,
	}
	return a.clnt
}

func DefaultAccount() *Account {
	accountConfig := currentConfig()
	if accountConfig.Default != nil {
		return accountConfig.Default
	}

	account := NewAccount()
	accountConfig.Default = account
	accountConfig.MergeAccounts(account)
	accountConfig.Save()
	return account
}

func AccountByScreenName(screenName string) *Account {
	accountConfig := currentConfig()
	for _, account := range accountConfig.Accounts {
		if account.ScreenName == screenName {
			return account
		}
	}

	account := NewAccount()
	accountConfig.MergeAccounts(account)
	accountConfig.Save()
	return account
}

func NewAccount() *Account {
	credential := auth.CredentialByClientName("Twitter for Android")
	accessToken := auth.Authenticate(credential)

	account := &Account{
		ConsumerKey:       credential.ConsumerKey,
		ConsumerSecret:    credential.ConsumerSecret,
		AccessToken:       accessToken.Token,
		AccessTokenSecret: accessToken.Secret,
		ScreenName:        accessToken.AdditionalData["screen_name"],
	}
	return account
}

func currentConfig() *AccountConfig {
	if fileExists(configFilePath()) {
		config := &AccountConfig{}
		data, err := ioutil.ReadFile(configFilePath())
		if err != nil {
			log.Fatal(err)
		}
		decoder := json.NewDecoder(bytes.NewReader(data))
		decoder.Decode(config)
		return config
	} else {
		emptyConfig := &AccountConfig{Default: nil, Accounts: []*Account{}}
		emptyConfig.Save()
		return emptyConfig
	}
}

func configFilePath() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	return usr.HomeDir + "/.thunderbolt"
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	return !os.IsNotExist(err)
}

type AccountConfig struct {
	Default  *Account
	Accounts []*Account
}

func (a *AccountConfig) Save() {
	configJson, err := json.Marshal(a)
	if err != nil {
		log.Fatal(err)
	}
	ioutil.WriteFile(configFilePath(), configJson, 0644)
}

func (a *AccountConfig) MergeAccounts(account *Account) {
	for _, currentAccount := range a.Accounts {
		if currentAccount.ScreenName == account.ScreenName {
			*currentAccount = *account
			return
		}
	}
	a.Accounts = append(a.Accounts, account)
}
