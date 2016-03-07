package main

import (
	"bufio"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"os"
)

const (
	FILENAME = "tweets"
)

func PostTweet(client *twitter.Client, tweet string) error {
	_, _, err := client.Statuses.Update(tweet, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	consumer_key := os.Getenv("TWITTER_CONS_KEY")
	consumer_secret := os.Getenv("TWITTER_CONS_SECRET")
	access_token := os.Getenv("TWITTER_ACCESS_TOKEN")
	access_secret := os.Getenv("TWITTER_ACCESS_SECRET")

	var err error
	var tweet string

	config := oauth1.NewConfig(consumer_key, consumer_secret)
	token := oauth1.NewToken(access_token, access_secret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	tweet, err = LoadTweet()
	if err != nil {
		fmt.Println(err)
	}
	err = CheckValid(tweet)
	if err != nil {
		fmt.Println(err)
		return
	}
	PostTweet(client, tweet)
}

func LoadTweet() (string, error) {
	f, _ := os.Open(FILENAME)
	reader := bufio.NewReader(f)
	tweet, isPrefix, err := reader.ReadLine()
	if isPrefix {
		fmt.Println("It's a prefex")
		return "", nil
	}
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	return string(tweet), nil
}

func CheckValid(tweet string) error {
	if len(tweet) > 140 {
		return fmt.Errorf("More than 140 chars")
	}
	return nil
}
