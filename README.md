# iot-to-pushover

## Setup Instructions

 * Create / configure 1click to SNS bridge first from https://github.com/aldrinleal/oneclick-bridge
 * Create a `.env` file with your `PUSHOVER_TOKEN` and `PUSHOVER_USER_KEY` into it
 * deploy (see below)
 * bind the SNS topic (from first step) into the new lambda (`i2p-dev-service`)

## Deploying

(requires make, go and yarn):

```
$ go get -v ./...
$ yarn install --frozen-lockfile
$ make && yarn sls deploy
```

## Details

