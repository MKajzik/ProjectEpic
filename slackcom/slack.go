package slackcom

import (
	"bytes"
	"errors"
	"net/http"
	"time"
)

//SendSlackMessage exported
func SendSlackMessage(webhookURL string, msg []byte) error {

	//create new POST request and add Header
	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(msg))
	if err != nil {
		return err
	}
	req.Header.Add("Content-type", "application/json")

	//Send request
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//read response and check errors
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}

	return nil
}
