package epic

import (
	"encoding/json"
	"fmt"
	"kazik/free/game/integration/slack"
	"kazik/free/game/rest"
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

	var freeGameobject FreeGame

	response, err := rest.SendGET(url)
	if err != nil {
		return FreeGame{}, err
	}

	if err := json.Unmarshal(*response, &freeGameobject); err != nil {
		return FreeGame{}, err
	}

	return freeGameobject, nil
}

func prepareJSON(freeGameObject FreeGame) ([]byte, string, error) {

	var msg slack.WebhookJSON

	text, num := checkFreeGame(freeGameObject)

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

	var gameNameObject FreeGame

	msg := CreateQuery(name)

	requestBody, err := json.Marshal(msg)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	response, err := rest.SendPOST("application/json", "https://www.epicgames.com/graphql", requestBody)
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	if err := json.Unmarshal(*response, &gameNameObject); err != nil {
		fmt.Println(err)
		return "", err
	}

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
