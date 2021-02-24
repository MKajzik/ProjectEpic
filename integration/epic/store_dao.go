package epic

import (
	"encoding/json"
	"fmt"
	"kazik/free/game/integration/slack"
	"kazik/free/game/rest"
	"time"
)

func checkFreeGame(freeGameobject FreeGame) ([]string, []int) {

	var games []string
	var num []int
	now := time.Now()

	for i := 0; i < len(freeGameobject.GetAllElements()); i++ {
		if len(freeGameobject.GetPromotionalOffers(i)) != 0 {
			if freeGameobject.GetPromotianlOfferStartDate(i).Before(now) || freeGameobject.GetPromotianlOfferEndDate(i).After(now) {
				if freeGameobject.GetPrice(i) == 0 {
					games = append(games, freeGameobject.GetElementTitle(i))
					num = append(num, i)
				}
			}
		}
	}

	if games == nil && num == nil {
		num = append(num, 400)
	}

	return games, num
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

	if num[0] == 400 {
		block := slack.CreateTextBlock("Dzisiaj nie ma zadnej gry do odebrania. Sorki :P!")
		msg.AddItem(block)

	} else {

		var images []string
		var urls []string

		for i := range text {
			image := searchForImage(num[i], freeGameObject)
			url, err := prepareURL(text[i])
			if err != nil {
				url = "https://www.epicgames.com/store/pl/free-games"
			}
			images = append(images, image)
			urls = append(urls, url)
		}

		block1 := slack.CreateTextBlock("Siema, dzisiaj Epic zaserwował nam nową dawkę darmowych gier. Poniżej sprawdźcie ją i nie zapomnijcie jej *ODEBRAĆ*!")
		msg.AddItem(block1)

		for i := range text {
			block1 := slack.CreateImageBlock(text[i], images[i])
			block2 := slack.CreateButtonBlock("Aby pobrać grę powyżej kliknij przycisk *ODBIERZ*", "ODBIERZ", urls[i])
			msg.AddItem(block1)
			msg.AddItem(block2)
		}
	}

	requestBody, err := json.Marshal(msg)
	if err != nil {
		return nil, "error", err
	}

	return requestBody, text[0], nil
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
		if gameNameObject.GetElementTitle(i) == name {
			return gameNameObject.GetProductSlug(i)
		}
	}
	return "fortnite"
}
