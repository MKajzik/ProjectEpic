package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	struktura "kazik/epic/darmowe/pkg"
)

func main() {

	jsonFile, err := os.Open("epic.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened test.json")

	byteValue, _ := ioutil.ReadAll(jsonFile)

	defer jsonFile.Close()

	var data struktura.Darmowe

	erro := json.Unmarshal(byteValue, &data)
	if erro != nil {
		fmt.Println(erro)
	}

	fmt.Println(data.Data.Catalog.SearchStore.Elements[1].Title)
}
