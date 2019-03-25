package twitter

import (
	"errors"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"net/url"
	"time"
)

type Twitter struct {
	client   *twitter.Client
	username string
}

const ExistTimeHour int = 24

func New(ck string, cs string, at string, atk string) (*Twitter, error) {
	config := oauth1.NewConfig(ck, cs)
	token := oauth1.NewToken(at, atk)
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client := twitter.NewClient(httpClient)
	fmt.Println(client.Trends.Available())
	user, _, err := client.Accounts.VerifyCredentials(&twitter.AccountVerifyParams{})
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, errors.New("there is no user")
	}
	return &Twitter{
		client:   client,
		username: user.ScreenName,
	}, nil
}

func (tw Twitter) GetMyTweet() ([]twitter.Tweet, error) {
	v := url.Values{}
	v.Set("screen_name", tw.username)
	tweets, _, err := tw.client.Timelines.UserTimeline(&twitter.UserTimelineParams{
		ScreenName: tw.username,
		Count:      200,
	})

	if err != nil {
		return nil, err
	}
	return tweets, nil
}

func (tw Twitter) Filtering(ts []twitter.Tweet) ([]twitter.Tweet, error) {

	var results []twitter.Tweet

	for _, t := range ts {
		tweetedTime, err := time.Parse("Mon Jan 02 15:04:05 -0700 2006", t.CreatedAt)
		if err != nil {
			return nil, err
		}
		d := time.Since(tweetedTime)
		if int(d.Hours()) >= ExistTimeHour {
			results = append(results, t)
		}
	}
	return results, nil
}

func (tw Twitter) DestroyTweets(ts []twitter.Tweet) error {
	for _, t := range ts {
		tw, _, err := tw.client.Statuses.Destroy(t.ID, &twitter.StatusDestroyParams{})
		fmt.Printf("deleted:\n %+v \n", tw.Text)
		if err != nil {
			return err
		}
	}
	return nil
}
