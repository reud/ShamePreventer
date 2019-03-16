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
	un := os.Getenv("TWITTER_USERNAME")
	bn := os.Getenv("BUCKET_NAME")
	wkr := strage.New(bn)
	wkr.DummyFunc()
	if ck == "" || cs == "" || at == "" || atk == "" {
		panic(errors.New("Something nil value"))

	}
	client := twitter.New(ck, cs, at, atk, un)
	tws, err := client.GetMyTweet()
	if err != nil {
		panic(err)
	}
	err = client.DestroyTweets(tws)
	if err != nil {
		panic(err)
	}
}
