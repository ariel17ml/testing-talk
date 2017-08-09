package main

import (
	"fmt"
	"testing"
)

func BenchmarkSample(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if x := fmt.Sprintf("%d", 42); x != "42" {
			b.Fatalf("Unexpected string: %s", x)
		}
	}
}

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib(b *testing.B) {
	b.Run("Fib3", func(b *testing.B) { benchmarkFib(3, b) })
	b.Run("Fib10", func(b *testing.B) { benchmarkFib(10, b) })
	b.Run("Fib40", func(b *testing.B) { benchmarkFib(40, b) })
}

func TestFailed(t *testing.T) {
	r := Fib(0)
	if r == 0 { // indeed
		t.Errorf("I didn't expect that: %v", r)
	}
}
