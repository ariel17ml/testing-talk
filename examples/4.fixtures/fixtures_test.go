package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"testing"
)

type Component struct {
	Id      int    `json:"id"`
	Address string `json:"address"`
}

type Model struct {
	Id         int         `json:"id"`
	Name       string      `json:"name"`
	Components []Component `json:"components"`
}

func TestWithFixtures(t *testing.T) {
	path := filepath.Join("testdata", "fixture.json")
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("Cannot read fixture file! %v", err)
	}

	var fixtures map[string][]Model
	if err = json.Unmarshal(dat, &fixtures); err != nil {
		t.Errorf("Cannot unmarshal fixtures! %v", err)
	}

	for _, m := range fixtures["objects"] {
		t.Log(fmt.Sprintf("#%v-%s@%s", m.Id, m.Name, m.Components[0].Address))
	}
}
