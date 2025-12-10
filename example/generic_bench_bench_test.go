package main

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func BenchmarkGenericAdd(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkGenericAdd[int](b, []testCaseBenchmarkGenericAdd[int]{
			// TODO: Add benchmark cases
		})
	})
}

func BenchmarkGenericAppend(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkGenericAppend[int](b, []testCaseBenchmarkGenericAppend[int]{
			// TODO: Add benchmark cases
		})
	})
}

func benchmarkGenericAdd[T constraints.Ordered](b *testing.B, tests []testCaseBenchmarkGenericAdd[T]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkGenericAdd[T]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkGenericAdd[T]) {}
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
				GenericAdd[T](
					bb.args.a,
					bb.args.b,
				)
			}
		})
	}
}

func benchmarkGenericAppend[T any](b *testing.B, tests []testCaseBenchmarkGenericAppend[T]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkGenericAppend[T]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkGenericAppend[T]) {}
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
				GenericAppend[T](
					bb.args.s,
					bb.args.v,
				)
			}
		})
	}
}

type testCaseBenchmarkGenericAdd[T constraints.Ordered] struct {
	name string
	args struct {
		a T
		b T
	}
	init    func(b *testing.B, tt *testCaseBenchmarkGenericAdd[T])
	cleanup func(b *testing.B, tt *testCaseBenchmarkGenericAdd[T])
}

type testCaseBenchmarkGenericAppend[T any] struct {
	name string
	args struct {
		s []T
		v T
	}
	init    func(b *testing.B, tt *testCaseBenchmarkGenericAppend[T])
	cleanup func(b *testing.B, tt *testCaseBenchmarkGenericAppend[T])
}
