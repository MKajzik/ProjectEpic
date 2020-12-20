package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	httpvalues "kazik/epic/darmowe/httpstructure"
	struktura "kazik/epic/darmowe/pkg"
)

const (
	slackTimeout time.Duration = 10 * time.Second
	token        string        = "xoxb-245805981380-1581140089558-iH4W0EqnKHcxdG6VY9w535yD"
)

func main() {

	var webhookURL string = "https://hooks.slack.com/services/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"
	var epicURL string = "https://store-site-backend-static.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL"

	done := make(chan bool, 1)
	go checkDailyGame(epicURL, webhookURL, done)

	<-done

}

func checkDailyGame(epicURL string, webhookURL string, done chan bool) {
	for {
		freeGame, err := GetEpicFreeGame(epicURL)
		if err != nil {
			fmt.Println(err)
			return
		}
		SendSlackMessage(webhookURL, freeGame)
		time.Sleep(30 * time.Second)
	}
	done <- true
}

//SendSlackMessage exported
func SendSlackMessage(webhookURL string, darmowe struktura.Darmowe) error {

	var image string
	text, num := checkFreeGame(darmowe)

	for i := 0; i < len(darmowe.Data.Catalog.SearchStore.Elements); i++ {
		if darmowe.Data.Catalog.SearchStore.Elements[num].KeyImages[i].Type != "VaultClosed" {
			image = darmowe.Data.Catalog.SearchStore.Elements[num].KeyImages[i].URL
		}
	}

	msg := httpvalues.Slack{}
	msg.Blocks = make([]httpvalues.Blocks, 2)

	msg.Blocks[0].Type = "section"
	msg.Blocks[0].Text.Type = "mrkdwn"
	msg.Blocks[0].Text.Text = "Siema, dzisiaj Epic zaserwował nam nową darmową grę. Poniżej sprawdźcie ją i nie zapomnijcie jej *ODEBRAĆ*!"
	msg.Blocks[1].Type = "section"
	msg.Blocks[1].Text.Type = "mrkdwn"
	msg.Blocks[1].Text.Text = text
	msg.Blocks[1].Accessory = &httpvalues.Accessory{}
	msg.Blocks[1].Accessory.Type = "image"
	msg.Blocks[1].Accessory.ImageURL = image
	msg.Blocks[1].Accessory.AltText = text

	//Encode by Marshal message to json
	requestBody, err := json.Marshal(msg)
	if err != nil {
		return err
	}
	fmt.Println(string(requestBody))

	//create new POST request and add Header
	req, err := http.NewRequest(http.MethodPost, webhookURL, bytes.NewBuffer(requestBody))
	if err != nil {
		return err
	}
	req.Header.Add("Content-type", "application/json")

	//Send request
	client := &http.Client{Timeout: slackTimeout}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	//read response and check errors
	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned from Slack")
	}

	return nil
}

//GetEpicFreeGame export
func GetEpicFreeGame(url string) (struktura.Darmowe, error) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)

	var darmowe struktura.Darmowe

	if err := json.Unmarshal(byteValue, &darmowe); err != nil {
		log.Fatalln(err)
	}

	err = errors.New("None of data match to new free game")
	return darmowe, nil

}

func checkFreeGame(darmowe struktura.Darmowe) (string, int) {

	now := time.Now()
	if darmowe.Data.Catalog.SearchStore.Elements[0].Promotions.PromotionalOffers[0].StartDate.Before(now) || darmowe.Data.Catalog.SearchStore.Elements[0].Promotions.PromotionalOffers[0].EndDate.After(now) {
		return darmowe.Data.Catalog.SearchStore.Elements[0].Title, 0
	}
	if darmowe.Data.Catalog.SearchStore.Elements[1].Promotions.PromotionalOffers[0].StartDate.Before(now) || darmowe.Data.Catalog.SearchStore.Elements[1].Promotions.PromotionalOffers[0].EndDate.After(now) {
		return darmowe.Data.Catalog.SearchStore.Elements[1].Title, 1
	}
	return "Brak darmowej gry", 400
}
