package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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

	freeGame, err := GetEpicFreeGame(epicURL)
	if err != nil {
		fmt.Println("Error in Function GetEpicFreeGame")
		return
	}

	SendSlackMessage(webhookURL, freeGame)

}

//SendSlackMessage exported
func SendSlackMessage(webhookURL string, msg string) error {

	//Encode by Marshal message to json
	requestBody, err := json.Marshal(httpvalues.Values{Text: msg})
	if err != nil {
		return err
	}

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
func GetEpicFreeGame(url string) (string, error) {

	now := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "http.GET ERROR", err
	}
	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)

	var darmowe struktura.Darmowe

	if err := json.Unmarshal(byteValue, &darmowe); err != nil {
		fmt.Println(err)
		return "json.Unmarshall ERROR", err
	}

	if darmowe.Data.Catalog.SearchStore.Elements[0].Promotions.PromotionalOffers[0].StartDate.Before(now) || darmowe.Data.Catalog.SearchStore.Elements[0].Promotions.PromotionalOffers[0].EndDate.After(now) {
		return darmowe.Data.Catalog.SearchStore.Elements[0].Title, nil
	}
	if darmowe.Data.Catalog.SearchStore.Elements[1].Promotions.PromotionalOffers[0].StartDate.Before(now) || darmowe.Data.Catalog.SearchStore.Elements[1].Promotions.PromotionalOffers[0].EndDate.After(now) {
		return darmowe.Data.Catalog.SearchStore.Elements[1].Title, nil
	}

	err = errors.New("None of data match to new free game")
	return "None", err

}
