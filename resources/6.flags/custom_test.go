package main

import (
	"flag"
	"testing"
)

var red, green, blue bool

func init() {
	flag.BoolVar(&red, "red", false, "run red tests")
	flag.BoolVar(&green, "green", false, "run green tests")
	flag.BoolVar(&blue, "blue", false, "run blue tests")
	flag.Parse()
}

func TestVerySlowMethod(t *testing.T) {
	t.Run("RedTest1", func(t *testing.T) {
		if !red {
			t.Skip()
		}

		t.Log("Ran red test 1")
	})

	t.Run("RedTest2", func(t *testing.T) {
		if !red {
			t.Skip()
		}

		t.Log("Ran red test 2")
	})

	t.Run("GreenTest1", func(t *testing.T) {
		if !green {
			t.Skip()
		}

		t.Log("Ran green test 1")
	})

	t.Run("GreenTest2", func(t *testing.T) {
		if !green {
			t.Skip()
		}

		t.Log("Ran green test 2")
	})

	t.Run("BlueTest1", func(t *testing.T) {
		if !blue {
			t.Skip()
		}

		t.Log("Ran blue test 1")
	})

}
