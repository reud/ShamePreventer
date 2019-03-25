package strage

import (
	"bytes"
	"cloud.google.com/go/storage"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

const TimeLayout = "20060102"

type Worker struct {
	bucketName string
	folderName string
}

func New(bucketName string) Worker {
	t := time.Now()
	return Worker{
		bucketName: bucketName,
		folderName: t.Format(TimeLayout),
	}
}

func (wo Worker) DummyFunc() {
	fmt.Println("Hello World!")
}

func (wo Worker) SaveTweets(tweets []twitter.Tweet) error {
	if len(tweets) == 0 {
		return nil
	}
	if err := wo.put(wo.folderName+"/tweets.txt", []byte(getTweetStrS(tweets))); err != nil {
		return err
	}
	for index, tweet := range tweets {
		if err := wo.savePhoto(fmt.Sprintf("photo%+v", index+1), tweet); err != nil {
			return err
		}
	}
	return nil
}

func (wo Worker) savePhoto(fileId string, tweet twitter.Tweet) error {
	if tweet.Entities.Media == nil {
		fmt.Println("This tweet had no content")
		return nil
	}
	for index, media := range tweet.Entities.Media {
		resp, err := http.Get(media.MediaURL)
		if err != nil {
			return err
		}
		byteArray, _ := ioutil.ReadAll(resp.Body)
		if err = wo.put(fmt.Sprintf("%+v/photo/%+v-%+v.jpg", wo.folderName, fileId, index+1), byteArray); err != nil {
			return err
		}

		if err = resp.Body.Close(); err != nil {
			return err
		}

	}
	return nil
}

func getTweetStr(tweet twitter.Tweet) string {
	writeStr := ""
	writeStr += "Tweet ID:" + strconv.FormatInt(tweet.ID, 10) + "\n"
	writeStr += "TimeStamp:" + tweet.CreatedAt + "\n"
	writeStr += "Source:" + tweet.Source + "\n"
	writeStr += "Text:" + tweet.Text + "\n"
	if tweet.RetweetedStatus != nil {
		writeStr += "RetweetedStatusId:" + strconv.FormatInt(tweet.RetweetedStatus.ID, 10) + "\n"
		writeStr += "RetweetedStatusUserId:" + strconv.FormatInt(tweet.RetweetedStatus.User.ID, 10) + "\n"
		writeStr += "RetweetedStatusTimeStamp:" + tweet.RetweetedStatus.CreatedAt + "\n"
	}
	if tweet.Entities.Urls != nil {
		for index, url := range tweet.Entities.Urls {
			writeStr += fmt.Sprintf("URL%+v:%+v\n", index+1, url)
		}
	}
	return writeStr
}

func getTweetStrS(tweets []twitter.Tweet) string {
	result := ""
	for _, tweet := range tweets {
		result += getTweetStr(tweet)
		result += "----------------------------------------------------------\n"
	}
	return result
}

func (wo Worker) put(path string, data []byte) error {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return err
	}

	w := client.Bucket(wo.bucketName).Object(path).NewWriter(ctx)

	if n, err := w.Write(data); err != nil {
		return err
	} else if n != len(data) {
		return err
	}
	if err := w.Close(); err != nil {
		return err
	}

	if err := client.Close(); err != nil {
		return err
	}
	return nil
}

func (wo Worker) Get(path string) ([]byte, error) {
	ctx := context.Background()
	client, err := storage.NewClient(ctx)
	if err != nil {
		return nil, err
	}

	r, err := client.Bucket(wo.bucketName).Object(path).NewReader(ctx)
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		return nil, err
	}
	if err := r.Close(); err != nil {
		return nil, err
	}

	if err := client.Close(); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
