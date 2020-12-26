package rest

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

//SendPOST export
func SendPOST(contentType string, URL string, msg []byte) (*[]byte, error) {

	response, err := http.Post(URL, contentType, bytes.NewBuffer(msg))
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("Bad Request StatusCode =/= 200")
	}

	byteValue, _ := ioutil.ReadAll(response.Body)

	return &byteValue, nil
}

//SendGET export
func SendGET(url string) (*[]byte, error) {

	response, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("Bad Request StatusCode =/= 200")
	}

	byteValue, _ := ioutil.ReadAll(response.Body)

	return &byteValue, nil
}
