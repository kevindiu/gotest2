package main

import (
	"testing"
)

func FuzzGenericSum(f *testing.F) {
	// TODO: Initialize concrete type for T
	type T = int
	// TODO: Add seed corpus
	f.Fuzz(func(t *testing.T, a T, b T) {
		GenericSum(
			a,
			b,
		)
	})
}

func FuzzSwap(f *testing.F) {
	// TODO: Initialize concrete type for T
	type T = int
	// TODO: Add seed corpus
	f.Fuzz(func(t *testing.T, a T, b T) {
		Swap(
			a,
			b,
		)
	})
}
