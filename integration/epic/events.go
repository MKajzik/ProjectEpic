package epic

import (
	"fmt"
	"kazik/free/game/integration/slack"
	"time"
)

//CheckDailyGame export
func CheckDailyGame(epicURL string, webhookURL string, done chan bool) {
	var previousGame string
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
		if previousGame != actualGame {
			slack.SendSlackMessage(webhookURL, epicJSON)
			previousGame = actualGame
		}
		time.Sleep(30 * time.Minute)
	}
	done <- true
}
