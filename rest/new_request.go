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
		eror := fmt.Sprintf("POST: Bad Request StatusCode = %d", response.StatusCode)
		return nil, errors.New(eror)
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
		eror := fmt.Sprintf("GET: Bad Request StatusCode = %d", response.StatusCode)
		return nil, errors.New(eror)
	}

	byteValue, _ := ioutil.ReadAll(response.Body)

	return &byteValue, nil
}
