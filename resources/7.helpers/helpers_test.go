package main

import (
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

var testFile string = "/tmp/dat1"

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Error(err)
	}
}

func testRequest(t *testing.T, url string) func(t *testing.T) {
	resp, err := http.Get(url)
	checkError(t, err)

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	checkError(t, err)

	err = ioutil.WriteFile(testFile, body, 0644)
	checkError(t, err)

	return func(t *testing.T) {
		err := os.Remove(testFile)
		checkError(t, err)
	}
}

func TestWithHelper(t *testing.T) {
	tearDown := testRequest(t, "http://www.google.com")
	defer tearDown(t)
}

func TestWithHelperShorter(t *testing.T) {
	defer testRequest(t, "http://www.example.com")(t)
}
