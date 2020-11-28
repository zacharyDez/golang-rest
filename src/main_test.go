package main

import (
	"io/ioutil"
	"net/http"
	"testing"
)

func TestHomePageHelloWorld(t *testing.T) {
	url := "http://localhost:9090"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		t.Errorf("Homepage returned error.")
	}

	res, err := client.Do(req)

	if res.StatusCode != 200 {
		t.Errorf("Homepage did not return status 200. Got: %d", res.StatusCode)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if string(body) != "Hello World!" {
		t.Errorf("Homepage did not return Hello World! Got: " + string(body))
	}
}

func TestHomePageName(t *testing.T) {
	url := "http://localhost:9090?name=Sarah"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		t.Errorf("Homepage returned error.")
	}

	res, err := client.Do(req)

	if res.StatusCode != 200 {
		t.Errorf("Homepage did not return status 200. Got: %d", res.StatusCode)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if string(body) != "Hello Sarah!" {
		t.Errorf("Homepage did not return Hello World! Got: " + string(body))
	}
}
