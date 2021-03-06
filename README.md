# iot-to-pushover

A bridge from SNS Events (containing 1click data) to [Pushover](https://pushover.net/)

![image](https://user-images.githubusercontent.com/253109/115102121-9bf7d100-9f0e-11eb-902b-57e1d7170c9f.png)

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

