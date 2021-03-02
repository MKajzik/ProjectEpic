package main

import (
	"kazik/free/game/integration/epic"
	"log"
	"os"
)

func main() {

	webhookURL, webhookBool := os.LookupEnv("SLACK_URL")
	epicURL, epicBool := os.LookupEnv("EPIC_URL")

	if webhookBool == false {
		log.Fatal("Unable to find Slack URL in environment variables!")
	}
	if epicBool == false {
		log.Fatal("Unable to find Epic URL in environment variables!")
	}

	done := make(chan bool, 1)
	go epic.CheckDailyGame(epicURL, webhookURL, done)

	<-done

}
