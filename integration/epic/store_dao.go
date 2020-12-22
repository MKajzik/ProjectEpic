package epic

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"kazik/free/game/integration/slack"
	"kazik/free/game/rest"
	"net/http"
	"time"
)

func checkFreeGame(freeGameobject FreeGame) (string, int) {

	now := time.Now()
	for i := 0; i < len(freeGameobject.Data.Catalog.SearchStore.Elements); i++ {
		if len(freeGameobject.Data.Catalog.SearchStore.Elements[i].Promotions.PromotionalOffers) != 0 {
			if freeGameobject.Data.Catalog.SearchStore.Elements[i].Promotions.PromotionalOffers[0].StartDate.Before(now) || freeGameobject.Data.Catalog.SearchStore.Elements[i].Promotions.PromotionalOffers[0].EndDate.After(now) {
				return freeGameobject.Data.Catalog.SearchStore.Elements[i].Title, i
			}
		}
	}
	return "Brak darmowej gry", 400
}

func getEpicFreeGame(url string) (FreeGame, error) {

	JSONBytes, err := rest.GetJSON(url)
	if err != nil {
		return FreeGame{}, err
	}

	var freeGameobject FreeGame

	if err := json.Unmarshal(*JSONBytes, &freeGameobject); err != nil {
		return FreeGame{}, err
	}

	err = errors.New("None of data match to new free game")
	return freeGameobject, nil

}

func prepareJSON(freeGameObject FreeGame) ([]byte, string, error) {

	text, num := checkFreeGame(freeGameObject)

	var msg slack.WebhookJSON

	if num == 400 {

		msg = slack.CreateNoGameBlock()

	} else {
		image := searchForImage(num, freeGameObject)
		url, err := prepareURL(text)
		if err != nil {
			url = "https://www.epicgames.com/store/pl/free-games"
		}

		msg = slack.CreateGameBlocks(text, image, url, true)

	}

	requestBody, err := json.Marshal(msg)
	if err != nil {
		return nil, "error", err
	}

	return requestBody, text, nil
}

func searchForImage(num int, freeGameobject FreeGame) string {
	var image string
	for i := 0; i < len(freeGameobject.Data.Catalog.SearchStore.Elements[num].KeyImages); i++ {
		if freeGameobject.Data.Catalog.SearchStore.Elements[num].KeyImages[i].Type != "VaultClosed" {
			image = freeGameobject.Data.Catalog.SearchStore.Elements[num].KeyImages[i].URL
		}
	}

	return image
}

func prepareURL(name string) (string, error) {

	//TODO Refactor this function

	msg := CreateQuery(name)

	requestBody, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("blad 1")
		return "error", err
	}

	fmt.Println(string(requestBody))

	response, err := http.Post("https://www.epicgames.com/graphql", "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("blad 3")
		return "", err
	}

	byteValue, _ := ioutil.ReadAll(response.Body)

	fmt.Println(string(byteValue))

	var gameNameObject FreeGame

	if err := json.Unmarshal(byteValue, &gameNameObject); err != nil {
		fmt.Println("blad 4")
		return "", err
	}
	fmt.Println(gameNameObject.Data.Catalog.SearchStore.Elements[0].Title)

	sufix := searchForURL(gameNameObject, name)

	URL := "https://www.epicgames.com/store/pl/product/" + sufix

	return URL, nil
}

func searchForURL(gameNameObject FreeGame, name string) string {

	for i := 0; i < len(gameNameObject.Data.Catalog.SearchStore.Elements); i++ {
		if gameNameObject.Data.Catalog.SearchStore.Elements[i].Title == name {
			return gameNameObject.Data.Catalog.SearchStore.Elements[i].ProductSlug
		}
	}
	return "fortnite"
}
