package main

import "testing"

func TestSuite(t *testing.T) {

	t.Run("useCase1", func(t *testing.T) {
		t.Log("Use case 1")
	})

	t.Run("useCase2", func(t *testing.T) {
		t.Log("Use case 2")
	})
}
