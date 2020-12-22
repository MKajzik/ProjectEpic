package main

import (
	epic "kazik/free/game/integration/epic"
)

func main() {

	var webhookURL string = "https://hooks.slack.com/services/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"
	var epicURL string = "https://store-site-backend-static.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL"

	done := make(chan bool, 1)
	go epic.CheckDailyGame(epicURL, webhookURL, done)

	<-done

}
