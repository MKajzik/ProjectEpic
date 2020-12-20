package epic

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"kazik/epic/darmowe/slackcom"
	"log"
	"net/http"
	"time"
)

func checkFreeGame(darmowe Darmowe) (string, int) {

	now := time.Now()
	for i := 0; i < len(darmowe.Data.Catalog.SearchStore.Elements); i++ {
		if len(darmowe.Data.Catalog.SearchStore.Elements[i].Promotions.PromotionalOffers) != 0 {
			if darmowe.Data.Catalog.SearchStore.Elements[i].Promotions.PromotionalOffers[0].StartDate.Before(now) || darmowe.Data.Catalog.SearchStore.Elements[i].Promotions.PromotionalOffers[0].EndDate.After(now) {
				return darmowe.Data.Catalog.SearchStore.Elements[i].Title, i
			}
		}

	}

	return "Brak darmowej gry", 400
}

//CheckDailyGame export
func CheckDailyGame(epicURL string, webhookURL string, done chan bool) {
	for {
		freeGame, err := getEpicFreeGame(epicURL)
		if err != nil {
			fmt.Println(err)
			return
		}
		epicJSON, err := prepareJSON(freeGame)
		if err != nil {
			fmt.Println(err)
			return
		}

		slackcom.SendSlackMessage(webhookURL, epicJSON)
		time.Sleep(30 * time.Second)
	}
	done <- true
}

func getEpicFreeGame(url string) (Darmowe, error) {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)

	var darmowe Darmowe

	if err := json.Unmarshal(byteValue, &darmowe); err != nil {
		log.Fatalln(err)
	}

	err = errors.New("None of data match to new free game")
	return darmowe, nil

}

func prepareJSON(darmowe Darmowe) ([]byte, error) {
	var image string
	text, num := checkFreeGame(darmowe)
	if num == 400 {
		return []byte("Dzisiaj nie ma zadnej gry do odebrania. Sorki :P"), nil
	}

	for i := 0; i < len(darmowe.Data.Catalog.SearchStore.Elements[num].KeyImages); i++ {
		if darmowe.Data.Catalog.SearchStore.Elements[num].KeyImages[i].Type != "VaultClosed" {
			image = darmowe.Data.Catalog.SearchStore.Elements[num].KeyImages[i].URL
		}
	}

	msg := slackcom.Slack{}
	msg.Blocks = make([]slackcom.Blocks, 2)

	msg.Blocks[0].Type = "section"
	msg.Blocks[0].Text.Type = "mrkdwn"
	msg.Blocks[0].Text.Text = "Siema, dzisiaj Epic zaserwował nam nową darmową grę. Poniżej sprawdźcie ją i nie zapomnijcie jej *ODEBRAĆ*!"
	msg.Blocks[1].Type = "section"
	msg.Blocks[1].Text.Type = "mrkdwn"
	msg.Blocks[1].Text.Text = text
	msg.Blocks[1].Accessory = &slackcom.Accessory{}
	msg.Blocks[1].Accessory.Type = "image"
	msg.Blocks[1].Accessory.ImageURL = image
	msg.Blocks[1].Accessory.AltText = text

	//Encode by Marshal message to json
	requestBody, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}
	return requestBody, nil
}
