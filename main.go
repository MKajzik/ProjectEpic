package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"

	httpvalues "kazik/epic/darmowe/httpstructure"
	struktura "kazik/epic/darmowe/pkg"
)

//DefaultSlackTimeout default
const DefaultSlackTimeout = 10 * time.Second

func main() {

	jsonFile, err := os.Open("epic.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened test.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	defer jsonFile.Close()

	var darmowe struktura.Darmowe

	//err := json.Unmarshal(byteValue, &data)
	if err := json.Unmarshal(byteValue, &darmowe); err != nil {
		fmt.Println(err)
	}

	//var token string = "xoxb-245805981380-1581140089558-iH4W0EqnKHcxdG6VY9w535yD"

	//fmt.Println(data.Data.Catalog.SearchStore.Elements[1].Title)

	var freeGame string = "Darmowa gra w EpicGames Store: " + darmowe.Data.Catalog.SearchStore.Elements[1].Title

	requestBody, err := json.Marshal(httpvalues.Values{Text: freeGame})
	if err != nil {
		fmt.Println(err)
	}

	req, err := http.NewRequest("POST", "https://hooks.slack.com/services/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", bytes.NewBuffer(requestBody))
	req.Header.Add("Content-type", "application/json")
	if err != nil {
		fmt.Println(err)
	}

	//Send request
	client := &http.Client{Timeout: DefaultSlackTimeout}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
	}

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)
	if buf.String() != "ok" {
		fmt.Println(buf.String())
	}
}
