package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/aldrinleal/iot-to-pushover/types"
	"github.com/aldrinleal/iot-to-pushover/util"
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/joomcode/errorx"
	"strings"

	log "github.com/sirupsen/logrus"

	"github.com/go-resty/resty/v2"
)

var sounds = map[string]string{
	"SINGLE": "vibrate",
	"DOUBLE": "tugboat",
	"LONG":   "siren",
}

func snsHandler(event events.SNSEvent) (err error) {
	outputBuf := bytes.NewBuffer([]byte{})

	err = json.NewEncoder(outputBuf).Encode(event)

	if nil != err {
		err = errorx.Decorate(err, "decoding buffer")

		log.Warnf("Oops: %s", err)

		return err
	}

	fmt.Println("payload: " + outputBuf.String())

	for _, snsEvent := range event.Records {
		parsedEvent := &types.OneClickEvent{}

		err = json.NewDecoder(strings.NewReader(snsEvent.SNS.Message)).Decode(parsedEvent)

		if nil != err {
			err = errorx.Decorate(err, "decoding message: '%s'", snsEvent.SNS.Message)

			log.Warnf("Oops: %s", err)

			return err
		}

		log.Infof("body: %+v", parsedEvent)

		err = reportEvent(parsedEvent)

		if nil != err {
			err = errorx.Decorate(err, "calling reportx")

			log.Warnf("Oops: %s", err)

			return err
		}
	}

	return nil
}

func reportEvent(event *types.OneClickEvent) (err error) {
	buttonClickedEvent := event.DeviceEvent.ButtonClicked

	client := resty.New()

	if nil != buttonClickedEvent {
		remainingLife, err := event.DeviceInfo.RemainingLife.Float64()

		if nil != err {
			return err
		}

		remainingLifeAsInt := int(remainingLife)

		clickType := buttonClickedEvent.ClickType

		soundToUse := sounds[clickType]

		title := fmt.Sprintf("Event: %s from %s (@ %d %%)", clickType, event.DeviceInfo.DeviceID, remainingLifeAsInt)

		meta := []string{}

		for k, v := range event.PlacementInfo.Attributes {
			meta = append(meta, fmt.Sprintf("%s: %s", k, v))
		}

		for k, v := range event.DeviceInfo.Attributes {
			meta = append(meta, fmt.Sprintf("%s: %s", k, v))
		}

		messageBody := strings.Join(meta, "\n")

		pushoverToken := util.EnvIf("PUSHOVER_TOKEN", "MISSING")
		pushoverUser := util.EnvIf("PUSHOVER_USER_KEY", "MISSING")
		response, err := client.R().
			SetHeaders(map[string]string{
				"Content-Type": "application/x-www-form-urlencoded",
			}).
			SetFormData(map[string]string{
				"token":    pushoverToken,
				"user":     pushoverUser,
				"title":    title,
				"message":  messageBody,
				"priority": "1",
				"expire":   "1800",
				"sound":    soundToUse,
			}).
			Post("https://api.pushover.net/1/messages.json")

		if nil != err {
			err = errorx.Decorate(err, "publishint to pushover")

			return err
		}

		log.Infof("Published: (%03d) (%s)", response.StatusCode(), response.String())
	}

	return nil
}

func main() {
	if util.IsRunningOnLambda() {
		lambda.Start(snsHandler)
	} else {
		// log.Fatalf("Oops", http.ListenAndServe(":"+util.EnvIf("PORT", "8000"), engine))
	}
}
