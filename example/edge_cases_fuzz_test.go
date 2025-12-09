package main

import (
	"testing"
)

func FuzzFuncFactory(f *testing.F) {
	// TODO: Add seed corpus
	f.Fuzz(func(t *testing.T, n int) {
		FuncFactory(
			n,
		)
	})
}
