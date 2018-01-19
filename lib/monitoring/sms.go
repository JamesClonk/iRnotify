package monitoring

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/JamesClonk/iRnotify/lib/env"
	"github.com/JamesClonk/iRnotify/lib/racers"
	cfenv "github.com/cloudfoundry-community/go-cfenv"
)

var clientId string

func init() {
	// check for VCAP_SERVICES first
	vcap, err := cfenv.Current()
	if err != nil {
		log.Println("Could not parse VCAP environment variables")
		log.Println(err)
	} else {
		service, err := vcap.Services.WithName("irnotifysms")
		if err != nil {
			log.Println("Could not find irnotifysms service in VCAP_SERVICES")
			log.Println(err)
		} else {
			clientId = fmt.Sprintf("%v", service.Credentials["client_id"])
		}
	}

	// if IRNOTIFY_CLIENT_ID is not yet set then try to read it from ENV
	if len(clientId) == 0 {
		clientId = env.MustGet("IRNOTIFY_CLIENT_ID")
	}
}

func notify(racer racers.Racer) error {
	// 0123456789:242335#38135;987654321:291357
	receivers := strings.SplitN(env.MustGet("IRNOTIFY_RECEIVERS"), ";", -1)

	for _, receiver := range receivers {
		fields := strings.SplitN(receiver, ":", 2)
		phoneNum := fields[0]
		watchedRacers := strings.SplitN(fields[1], "#", -1)

		for _, watchedID := range watchedRacers {
			if watchedID == fmt.Sprintf("%d", racer.CustID) {
				text := fmt.Sprintf("%s is in a %s session right now, driving a %s on %s",
					racer.Name, racer.SessionTypeName, racer.CarName, racer.TrackName)
				if err := sendSms(phoneNum, text); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func sendSms(phoneNum, text string) error {
	var jsonStr = []byte(`{"to": "{` + phoneNum + `}", "text": "` + text + `"}`)
	req, err := http.NewRequest("POST", "https://api.swisscom.com/messaging/sms", bytes.NewBuffer(jsonStr))

	req.Header.Set("SCS-Version", "2")
	req.Header.Set("client_id", clientId)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	log.Printf("Sending notification SMS to [%s]: %s\n", phoneNum, text)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusNoContent {
		return fmt.Errorf("SMS status code: %v", resp.StatusCode)
	}
	return nil
}
