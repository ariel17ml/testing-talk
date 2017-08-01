package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"path/filepath"
	"strings"
	"testing"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	path := filepath.Join("testdata", "output.json")
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, string(dat))
}

func TestGolden(t *testing.T) {
	req, _ := http.NewRequest("GET", "/", nil)
	w := httptest.NewRecorder()

	handler := http.HandlerFunc(Handler)
	handler.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Status code must be 200; was %v", w.Code)
	}

	path := filepath.Join("testdata", "output.golden")
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("Failed to fetch content from golden file! %v", err)
	}

	body := w.Body.String()
	if !strings.Contains(string(dat), body) {
		t.Errorf("Golden file content is not in HTTP response! %v", body)
	}
}
