package main

import (
	"bytes"
	"encoding/json"
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

func DefaultAccount() *Account {
	accountConfig := currentConfig()
	if accountConfig.Default != nil {
		return accountConfig.Default
	}

	credential := auth.CredentialByClientName("Twitter for Android")
	accessToken := auth.Authenticate(credential)

	account := &Account{
		ConsumerKey:       credential.ConsumerKey,
		ConsumerSecret:    credential.ConsumerSecret,
		AccessToken:       accessToken.Token,
		AccessTokenSecret: accessToken.Secret,
		ScreenName:        accessToken.AdditionalData["screen_name"],
	}
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

	credential := auth.CredentialByClientName("Twitter for Android")
	accessToken := auth.Authenticate(credential)

	account := &Account{
		ConsumerKey:       credential.ConsumerKey,
		ConsumerSecret:    credential.ConsumerSecret,
		AccessToken:       accessToken.Token,
		AccessTokenSecret: accessToken.Secret,
		ScreenName:        accessToken.AdditionalData["screen_name"],
	}
	accountConfig.MergeAccounts(account)
	accountConfig.Save()
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
