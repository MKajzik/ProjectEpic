package rest

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

//CreateRequest export
func createRequest(httpMethod string, URL string, msg []byte) (*http.Request, error) {

	req, err := http.NewRequest(httpMethod, URL, bytes.NewBuffer(msg))
	if err != nil {
		return nil, err
	}
	if httpMethod == "GET" {
		req.Header.Add("Accept", "application/json")
	} else if httpMethod == "POST" {
		req.Header.Add("Content-type", "application/json")
	}
	return req, nil
}

//DoRequest export
func doRequest(request *http.Request) (*http.Response, error) {

	client := &http.Client{Timeout: 10 * time.Second}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	return response, nil
}

//CheckResponse export
func checkResponse(response *http.Response) error {

	buf := new(bytes.Buffer)
	buf.ReadFrom(response.Body)
	if buf.String() != "ok" {
		return errors.New("Non-ok response returned")
	}
	return nil
}

//PrepareAndExecuteRequest export
func PrepareAndExecuteRequest(httpMethod string, URL string, msg []byte) (*http.Response, error) {
	request, err := createRequest(httpMethod, URL, msg)
	if err != nil {
		return nil, err
	}
	response, err := doRequest(request)
	if err != nil {
		return nil, err
	}
	err = checkResponse(response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

//GetJSON export
func GetJSON(url string) (*[]byte, error) {

	response, err := http.Get("https://store-site-backend-static.ak.epicgames.com/freeGamesPromotions?locale=pl&country=PL&allowCountries=PL")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, err
	}

	//checkResponse(response)

	byteValue, _ := ioutil.ReadAll(response.Body)

	return &byteValue, nil
}
