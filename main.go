package main

import (
	"errors"
	"fmt"
	"github.com/reud/ShamePreventer/strage"
	"github.com/reud/ShamePreventer/twitter"
	"os"
)

func main() {
	fmt.Println(os.Environ())
	ck := os.Getenv("TWITTER_CONSUMER_KEY")
	cs := os.Getenv("TWITTER_CONSUMER_SECRET")
	at := os.Getenv("TWITTER_ACCESS_TOKEN")
	atk := os.Getenv("TWITTER_ACCESS_TOKEN_SECRET")
	bn := os.Getenv("BUCKET_NAME")
	wkr := strage.New(bn)
	if ck == "" || cs == "" || at == "" || atk == "" {
		panic(errors.New("Something had nil value"))
	}
	client, err := twitter.New(ck, cs, at, atk)
	if err != nil {
		panic(err)
	}
	tws, err := client.GetMyTweet()
	if err != nil {
		panic(err)
	}
	if err := wkr.SaveTweets(tws); err != nil {
		panic(err)
	}
	err = client.DestroyTweets(tws)
	if err != nil {
		panic(err)
	}
}
