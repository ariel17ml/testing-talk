package main

import (
	"flag"
	"testing"
)

var acceptance string

func init() {
	flag.StringVar(&acceptance, "acceptance", "", "run the acceptance tests")
	flag.Parse()
}

func TestVerySlowMethod(t *testing.T) {
	if testing.Short() {
		t.Skip("this test must be skipped in short mode.")
	}
}
