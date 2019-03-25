Shame Preventer v1.1
====

Tweet freely, like instagram stories

updated! if you want to use before version, watch (it)[https://github.com/reud/ShamePreventer/tree/v1.0-master]

## Description

(v1.1 new feature! this steps will apply if tweet after 24 hours)

1 It archives your tweets (at GCP)

2 Then, deletes your tweets

If you set cron working every day, It will be like Instagram Stories.

## VS. 

Instagtam Stories: They seem happy, so I feel like a LOSER.

Twitter+Shame Preventer: **COMFORTABLE**

## Requirement

see [go.mod](https://github.com/reud/ShamePreventer/blob/master/go.mod)
## Usage

Before start using, you have to get key.json from GCP, and put it on project directory.

```bash
$ git clone https://github.com/reud/ShamePreventer.git
$ cd ShamePreventer
$ docker build -t ShamePreventer . --build-arg GOARCH="amd64" 
# if you want to use on rasberry Pi, use instead of GOARCH="arm"
$ docker run -e TWITTER_CONSUMER_KEY="YOUR_KEY" \
  -e TWITTER_CONSUMER_SECRET="YOUR_SECRET" \
  -e TWITTER_ACCESS_TOKEN="YOUR_TOKEN" \
  -e TWITTER_ACCESS_TOKEN_SECRET="YOUR_TOKREN_SECRET" \
  -e BUCKET_NAME="YOUR_GCP_BUCKET_NAME" ShamePreventer
```


## Licence

[MIT](https://github.com/reud/MIT_LICENSE/blob/master/README.md)

## Author

[reud](https://github.com/reud)