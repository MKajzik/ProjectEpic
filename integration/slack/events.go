package slack

import (
	"kazik/free/game/rest"
)

//SendSlackMessage exported
func SendSlackMessage(webhookURL string, msg []byte) error {

	_, err := rest.SendPOST("application/json", webhookURL, msg)
	if err != nil {
		return err
	}

	return nil
}
