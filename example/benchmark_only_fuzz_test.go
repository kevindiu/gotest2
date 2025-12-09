package main

import (
	"testing"
)

func FuzzMultiply(f *testing.F) {
	// TODO: Add seed corpus
	f.Fuzz(func(t *testing.T, a int, b int) {
		Multiply(
			a,
			b,
		)
	})
}
