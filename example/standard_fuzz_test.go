package main

import (
	"testing"
)

func FuzzAdd(f *testing.F) {
	// TODO: Add seed corpus
	f.Fuzz(func(t *testing.T, a int, b int) {
		Add(
			a,
			b,
		)
	})
}

func FuzzDivMod(f *testing.F) {
	// TODO: Add seed corpus
	f.Fuzz(func(t *testing.T, a int, b int) {
		DivMod(
			a,
			b,
		)
	})
}
