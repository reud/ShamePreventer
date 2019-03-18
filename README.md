Shame Preventer
====

Tweet feel free, like instagram stories

## Description

1. Your tweets archiving (at GCP)
2. Delete your tweets 

If you set cron and working every day, It is like Instagram Stories.

## VS. 

Instagtam Stories: They seems to be happy, So I feel be a LOSER.

Twitter+Shame Preventer: **COMFORTABLE**

## Requirement

see [go.mod](https://github.com/reud/ShamePreventer/blob/master/go.mod)
## Usage

Before, You have to get key.json from GCP, and put on project directory.

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