package main

import (
	"testing"
)

func BenchmarkBox_Get(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkBox_Get[int](b, []testCaseBenchmarkBox_Get[int]{
			// TODO: Add benchmark cases
		})
	})
}

func BenchmarkBox_Put(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkBox_Put[int](b, []testCaseBenchmarkBox_Put[int]{
			// TODO: Add benchmark cases
		})
	})
}

func BenchmarkProcessContainer(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkProcessContainer[int](b, []testCaseBenchmarkProcessContainer[int]{
			// TODO: Add benchmark cases
		})
	})
}

func benchmarkBox_Get[T any](b *testing.B, tests []testCaseBenchmarkBox_Get[T]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkBox_Get[T]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkBox_Get[T]) {}
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
				bb.receiver.Get()
			}
		})
	}
}

func benchmarkBox_Put[T any](b *testing.B, tests []testCaseBenchmarkBox_Put[T]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkBox_Put[T]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkBox_Put[T]) {}
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
				bb.receiver.Put(
					bb.args.v,
				)
			}
		})
	}
}

func benchmarkProcessContainer[T any](b *testing.B, tests []testCaseBenchmarkProcessContainer[T]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkProcessContainer[T]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkProcessContainer[T]) {}
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
				ProcessContainer[T](
					bb.args.c,
					bb.args.val,
				)
			}
		})
	}
}

type testCaseBenchmarkBox_Get[T any] struct {
	name string
	args struct {
	}
	receiver *Box[T]
	init     func(b *testing.B, tt *testCaseBenchmarkBox_Get[T])
	cleanup  func(b *testing.B, tt *testCaseBenchmarkBox_Get[T])
}

type testCaseBenchmarkBox_Put[T any] struct {
	name string
	args struct {
		v T
	}
	receiver *Box[T]
	init     func(b *testing.B, tt *testCaseBenchmarkBox_Put[T])
	cleanup  func(b *testing.B, tt *testCaseBenchmarkBox_Put[T])
}

type testCaseBenchmarkProcessContainer[T any] struct {
	name string
	args struct {
		c   Container[T]
		val T
	}
	init    func(b *testing.B, tt *testCaseBenchmarkProcessContainer[T])
	cleanup func(b *testing.B, tt *testCaseBenchmarkProcessContainer[T])
}
