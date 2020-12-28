package main

import "kazik/free/game/integration/epic"

func main() {

	webhookURL := readFromConfigurationFile("webhookURL")
	epicURL := readFromConfigurationFile("epicURL")

	done := make(chan bool, 1)
	go epic.CheckDailyGame(epicURL, webhookURL, done)

	<-done

}
