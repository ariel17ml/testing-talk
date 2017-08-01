package main

import (
	"net/http"
	"testing"
	"time"
)

func TestSlow(t *testing.T) {

	if testing.Short() {
		t.Skip("Skiping this slow test")
	}

	for i := 0; i < 5; i++ {
		resp, err := http.Get("http://google.com/")
		t.Logf("Response: %v, error: %v", resp, err)
		time.Sleep(1000 * time.Millisecond)
	}
}

func TestQuick(t *testing.T) {
	http.Get("http://google.com/")
	t.Log("Running always")
}
