package service

import (
	"context"
	"testing"

	"github.com/kevindiu/gotest2/example/model"
	"github.com/kevindiu/gotest2/example/repository"
)

func BenchmarkNewBookService(b *testing.B) {
	type args struct {
		repo repository.Repository[model.Book, string]
	}
	type test struct {
		name    string
		args    args
		init    func(b *testing.B, tt *test)
		cleanup func(b *testing.B, tt *test)
	}
	tests := []test{
		// TODO: Add benchmark cases
	}
	defaultInit := func(b *testing.B, tt *test) {}
	defaultCleanup := func(b *testing.B, tt *test) {}
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
			for i := 0; i < b.N; i++ {
				NewBookService(
					bb.args.repo,
				)
			}
		})
	}
}

func BenchmarkBookService_CreateBook(b *testing.B) {
	type args struct {
		title  string
		author string
		isbn   string
	}
	type test struct {
		name     string
		args     args
		receiver *BookService
		init     func(b *testing.B, tt *test)
		cleanup  func(b *testing.B, tt *test)
	}
	tests := []test{
		// TODO: Add benchmark cases
	}
	defaultInit := func(b *testing.B, tt *test) {}
	defaultCleanup := func(b *testing.B, tt *test) {}
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
			for i := 0; i < b.N; i++ {
				bb.receiver.CreateBook(
					bb.args.title,
					bb.args.author,
					bb.args.isbn,
				)
			}
		})
	}
}

func BenchmarkBookService_GetBook(b *testing.B) {
	type args struct {
		id string
	}
	type test struct {
		name     string
		args     args
		receiver *BookService
		init     func(b *testing.B, tt *test)
		cleanup  func(b *testing.B, tt *test)
	}
	tests := []test{
		// TODO: Add benchmark cases
	}
	defaultInit := func(b *testing.B, tt *test) {}
	defaultCleanup := func(b *testing.B, tt *test) {}
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
			for i := 0; i < b.N; i++ {
				bb.receiver.GetBook(
					bb.args.id,
				)
			}
		})
	}
}

func BenchmarkBookService_ListBooks(b *testing.B) {
	type args struct {
	}
	type test struct {
		name     string
		args     args
		receiver *BookService
		init     func(b *testing.B, tt *test)
		cleanup  func(b *testing.B, tt *test)
	}
	tests := []test{
		// TODO: Add benchmark cases
	}
	defaultInit := func(b *testing.B, tt *test) {}
	defaultCleanup := func(b *testing.B, tt *test) {}
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
			for i := 0; i < b.N; i++ {
				bb.receiver.ListBooks()
			}
		})
	}
}

func BenchmarkBookService_BatchCreate(b *testing.B) {
	type args struct {
		ctx   context.Context
		books []model.Book
	}
	type test struct {
		name     string
		args     args
		receiver *BookService
		init     func(b *testing.B, tt *test)
		cleanup  func(b *testing.B, tt *test)
	}
	tests := []test{
		// TODO: Add benchmark cases
	}
	defaultInit := func(b *testing.B, tt *test) {}
	defaultCleanup := func(b *testing.B, tt *test) {}
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
			for i := 0; i < b.N; i++ {
				bb.receiver.BatchCreate(
					bb.args.ctx,
					bb.args.books,
				)
			}
		})
	}
}

func BenchmarkBookService_validateBook(b *testing.B) {
	type args struct {
		b model.Book
	}
	type test struct {
		name     string
		args     args
		receiver *BookService
		init     func(b *testing.B, tt *test)
		cleanup  func(b *testing.B, tt *test)
	}
	tests := []test{
		// TODO: Add benchmark cases
	}
	defaultInit := func(b *testing.B, tt *test) {}
	defaultCleanup := func(b *testing.B, tt *test) {}
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
			for i := 0; i < b.N; i++ {
				bb.receiver.validateBook(
					bb.args.b,
				)
			}
		})
	}
}
