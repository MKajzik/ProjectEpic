package main

import (
	epic "kazik/epic/darmowe/pkg"
)

const (
	token string = "xoxb-245805981380-1581140089558-iH4W0EqnKHcxdG6VY9w535yD"
)

func main() {

	var webhookURL string = "https://hooks.slack.com/services/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"
	var epicURL string = "https://store-site-backend-static.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL"

	done := make(chan bool, 2)
	go epic.CheckDailyGame(epicURL, webhookURL, done)

	<-done

}
