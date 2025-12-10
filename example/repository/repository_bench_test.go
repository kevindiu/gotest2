package repository

import (
	"testing"
)

func BenchmarkMemoryRepository_Create(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkMemoryRepository_Create[int, int](b, []testCaseBenchmarkMemoryRepository_Create[int, int]{
			// TODO: Add benchmark cases
		})
	})
}

func BenchmarkMemoryRepository_Get(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkMemoryRepository_Get[int, int](b, []testCaseBenchmarkMemoryRepository_Get[int, int]{
			// TODO: Add benchmark cases
		})
	})
}

func BenchmarkMemoryRepository_List(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkMemoryRepository_List[int, int](b, []testCaseBenchmarkMemoryRepository_List[int, int]{
			// TODO: Add benchmark cases
		})
	})
}

func BenchmarkMemoryRepository_Delete(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkMemoryRepository_Delete[int, int](b, []testCaseBenchmarkMemoryRepository_Delete[int, int]{
			// TODO: Add benchmark cases
		})
	})
}

func BenchmarkNewMemoryRepository(b *testing.B) {
	b.Run("int", func(b *testing.B) {
		benchmarkNewMemoryRepository[int, int](b, []testCaseBenchmarkNewMemoryRepository[int, int]{
			// TODO: Add benchmark cases
		})
	})
}

func benchmarkMemoryRepository_Create[T any, ID comparable](b *testing.B, tests []testCaseBenchmarkMemoryRepository_Create[T, ID]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_Create[T, ID]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_Create[T, ID]) {}
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
				bb.receiver.Create(
					bb.args.id,
					bb.args.entity,
				)
			}
		})
	}
}

func benchmarkMemoryRepository_Get[T any, ID comparable](b *testing.B, tests []testCaseBenchmarkMemoryRepository_Get[T, ID]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_Get[T, ID]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_Get[T, ID]) {}
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
				bb.receiver.Get(
					bb.args.id,
				)
			}
		})
	}
}

func benchmarkMemoryRepository_List[T any, ID comparable](b *testing.B, tests []testCaseBenchmarkMemoryRepository_List[T, ID]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_List[T, ID]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_List[T, ID]) {}
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
				bb.receiver.List()
			}
		})
	}
}

func benchmarkMemoryRepository_Delete[T any, ID comparable](b *testing.B, tests []testCaseBenchmarkMemoryRepository_Delete[T, ID]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_Delete[T, ID]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_Delete[T, ID]) {}
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
				bb.receiver.Delete(
					bb.args.id,
				)
			}
		})
	}
}

func benchmarkNewMemoryRepository[T any, ID comparable](b *testing.B, tests []testCaseBenchmarkNewMemoryRepository[T, ID]) {
	defaultInit := func(b *testing.B, tt *testCaseBenchmarkNewMemoryRepository[T, ID]) {}
	defaultCleanup := func(b *testing.B, tt *testCaseBenchmarkNewMemoryRepository[T, ID]) {}
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
				NewMemoryRepository[T, ID]()
			}
		})
	}
}

type testCaseBenchmarkMemoryRepository_Create[T any, ID comparable] struct {
	name string
	args struct {
		id     ID
		entity T
	}
	receiver *MemoryRepository[T, ID]
	init     func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_Create[T, ID])
	cleanup  func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_Create[T, ID])
}

type testCaseBenchmarkMemoryRepository_Get[T any, ID comparable] struct {
	name string
	args struct {
		id ID
	}
	receiver *MemoryRepository[T, ID]
	init     func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_Get[T, ID])
	cleanup  func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_Get[T, ID])
}

type testCaseBenchmarkMemoryRepository_List[T any, ID comparable] struct {
	name string
	args struct {
	}
	receiver *MemoryRepository[T, ID]
	init     func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_List[T, ID])
	cleanup  func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_List[T, ID])
}

type testCaseBenchmarkMemoryRepository_Delete[T any, ID comparable] struct {
	name string
	args struct {
		id ID
	}
	receiver *MemoryRepository[T, ID]
	init     func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_Delete[T, ID])
	cleanup  func(b *testing.B, tt *testCaseBenchmarkMemoryRepository_Delete[T, ID])
}

type testCaseBenchmarkNewMemoryRepository[T any, ID comparable] struct {
	name string
	args struct {
	}
	init    func(b *testing.B, tt *testCaseBenchmarkNewMemoryRepository[T, ID])
	cleanup func(b *testing.B, tt *testCaseBenchmarkNewMemoryRepository[T, ID])
}
