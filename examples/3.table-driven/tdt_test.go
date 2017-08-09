package main

import (
	"fmt"
	"testing"
)

func Formatter(in1, in2 string) string {
	return fmt.Sprintf("%s-xx-%s", in1, in2)
}

func TestFormatter(t *testing.T) {
	var table = map[string]struct{ A, B, Expected string }{
		"only-strings": {"a", "b", "a-xx-b"},
		"with-numbers": {"01", "02", "01-xx-02"},
		"failing-case": {"01", "02", "BBBBBBBB"},
	}

	for name, tc := range table {
		s := Formatter(tc.A, tc.B)
		if s != tc.Expected {
			t.Errorf("%s @ Formatter(%q, %q) => %q, want %q", name, tc.A, tc.B, s, tc.Expected)
		}
	}
}
