package main

import "golang.org/x/exp/constraints"

// GenericAdd adds two numbers.
func GenericAdd[T constraints.Ordered](a, b T) T {
	return a + b
}

// GenericAppend appends a value to a slice.
func GenericAppend[T any](s []T, v T) []T {
	return append(s, v)
}
