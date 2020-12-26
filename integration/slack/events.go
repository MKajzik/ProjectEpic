package slack

import (
	"fmt"
	"kazik/free/game/rest"
)

//SendSlackMessage exported
func SendSlackMessage(webhookURL string, msg []byte) error {

	_, err := rest.SendPOST("application/json", webhookURL, msg)
	fmt.Println(err)
	if err != nil {
		return err
	}

	return nil
}
