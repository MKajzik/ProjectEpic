package epic

import (
	"fmt"
	"kazik/free/game/integration/slack"
	"time"
)

//CheckDailyGame export
func CheckDailyGame(epicURL string, webhookURL string, done chan bool) {
	var previousGame []string

	for {
		freeGame, err := getEpicFreeGame(epicURL)
		if err != nil {
			fmt.Println(err)
			time.Sleep(30 * time.Second)
			continue
		}
		epicJSON, actualGame, err := prepareJSON(freeGame)
		if err != nil {
			fmt.Println(err)
			time.Sleep(30 * time.Second)
			continue
		}
		if len(previousGame) == 0 {
			slack.SendSlackMessage(webhookURL, epicJSON)
			previousGame = actualGame

		} else {
			if len(actualGame) < len(previousGame) {
				for i := range actualGame {
					if actualGame[i] != previousGame[i] {
						slack.SendSlackMessage(webhookURL, epicJSON)
						previousGame = actualGame
						break
					}
				}
			} else {
				for i := range previousGame {
					if previousGame[i] != actualGame[i] {
						slack.SendSlackMessage(webhookURL, epicJSON)
						previousGame = actualGame
						break
					}
				}
			}
		}
		time.Sleep(30 * time.Minute)
	}
	done <- true
}
