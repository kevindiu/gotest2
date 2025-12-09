package main

import (
	"testing"
)

func BenchmarkGenericMultiError(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkGenericMultiError[int](b, []testCaseBenchmarkGenericMultiError[int]{
			// TODO: Add benchmark cases
		})
	})
}

func benchmarkGenericMultiError[T any](b *testing.B, tests []testCaseBenchmarkGenericMultiError[T]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkGenericMultiError[T]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkGenericMultiError[T]) {}
	for _, bb := range tests {
		b.Run(bb.name, func(b *testing.B) {
			if bb.init == nil {
				bb.init = defaultInit
			}
			bb.init(b, &bb)
			if bb.cleanup == nil {
				bb.cleanup = defaultCleanup
			}
			defer bb.cleanup(b, &bb)
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				GenericMultiError(
					bb.args.val,
					bb.args.n,
				)
			}
		})
	}
}

type testCaseBenchmarkGenericMultiError[T any] struct {
	name string
	args struct {
		val T
		n   int
	}
	init    func(b *testing.B, tt *testCaseBenchmarkGenericMultiError[T])
	cleanup func(b *testing.B, tt *testCaseBenchmarkGenericMultiError[T])
}
