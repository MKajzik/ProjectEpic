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

	msg := slack.NewRequest()

	text, num := checkFreeGame(freeGameObject)

	if num == 400 {
		textBuilder := slack.NewTextBuilder()
		text1 := textBuilder.
			SetType("mrkdwn").
			SetText("Dzisiaj nie ma zadnej gry do odebrania. Sorki :P!").
			Build()

		blockBuilder := slack.NewBlockBuilder()
		block1 := blockBuilder.
			SetType("section").
			SetText(text1).
			Build()

		msg.AddItem(block1)

	} else {
		image := searchForImage(num, freeGameObject)
		url, err := prepareURL(text)
		if err != nil {
			url = "https://www.epicgames.com/store/pl/free-games"
		}
		textBuilder := slack.NewTextBuilder()
		accessoryBuilder := slack.NewAccessoryBuilder()
		blockBuilder := slack.NewBlockBuilder()

		block1 := blockBuilder.
			SetType("section").
			SetText(textBuilder.
				SetType("mrkdwn").
				SetText("Siema, dzisiaj Epic zaserwował nam nową darmową grę. Poniżej sprawdźcie ją i nie zapomnijcie jej *ODEBRAĆ*!").
				Build()).
			Build()

		textBuilder.Reset()
		accessoryBuilder.Reset()
		blockBuilder.Reset()

		block2 := blockBuilder.
			SetType("section").
			SetText(textBuilder.
				SetType("mrkdwn").
				SetText(text).
				Build()).
			SetAccessory(accessoryBuilder.
				SetType("image").
				SetImageURL(image).
				SetAltText(text).
				Build()).
			Build()

		textBuilder.Reset()
		accessoryBuilder.Reset()
		blockBuilder.Reset()

		block3 := blockBuilder.
			SetType("section").
			SetText(textBuilder.
				SetType("mrkdwn").
				SetText("Odbierz mnie pliska. *NO PLISKA*").
				Build()).
			SetAccessory(accessoryBuilder.
				SetType("button").
				SetText(textBuilder.
					SetType("plain_text").
					SetText("ODBIERZ").
					SetEmoji(true).
					Build()).
				SetValue("click_me_123").
				SetURL(url).
				SetActionID("button-action").
				Build()).
			Build()

		msg.AddItem(block1)
		msg.AddItem(block2)
		msg.AddItem(block3)
	}

	requestBody, err := json.Marshal(msg)
	if err != nil {
		return nil, "error", err
	}

	return requestBody, text, nil
}

func searchForImage(num int, freeGameobject FreeGame) string {
	var image string

	for i := 0; i < len(freeGameobject.GetAllKeyImages(num)); i++ {
		if freeGameobject.GetKeyImageType(num, i) != "VaultClosed" {
			image = freeGameobject.GetKeyImageURL(num, i)
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

	for i := 0; i < len(gameNameObject.GetAllElements()); i++ {
		if gameNameObject.GetTitle(i) == name {
			return gameNameObject.GetProductSlug(i)
		}
	}
	return "fortnite"
}
