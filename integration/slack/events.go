package slack

import (
	"fmt"
	"kazik/free/game/rest"
)

//SendSlackMessage exported
func SendSlackMessage(webhookURL string, msg []byte) error {

	fmt.Println("Slack")
	_, err := rest.SendPOST("application/json", webhookURL, msg)
	if err != nil {
		return err
	}

	return nil
}
