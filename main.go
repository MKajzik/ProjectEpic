package main

import (
	"kazik/free/game/integration/epic"
	"os"
)

func main() {

	webhookURL := os.Getenv("SLACK_URL")
	epicURL := os.Getenv("EPIC_URL")

	done := make(chan bool, 1)
	go epic.CheckDailyGame(epicURL, webhookURL, done)

	<-done

}
