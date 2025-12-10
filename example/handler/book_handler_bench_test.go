package handler

import (
	"net/http"
	"testing"

	"github.com/kevindiu/gotest2/example/service"
)

func BenchmarkNewBookHandler(b *testing.B) {
	type args struct {
		svc *service.BookService
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
				NewBookHandler(
					bb.args.svc,
				)
			}
		})
	}
}

func BenchmarkBookHandler_CreateBookHandler(b *testing.B) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type test struct {
		name     string
		args     args
		receiver *BookHandler
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
				bb.receiver.CreateBookHandler(
					bb.args.w,
					bb.args.r,
				)
			}
		})
	}
}

func BenchmarkBookHandler_GetBookHandler(b *testing.B) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type test struct {
		name     string
		args     args
		receiver *BookHandler
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
				bb.receiver.GetBookHandler(
					bb.args.w,
					bb.args.r,
				)
			}
		})
	}
}

func BenchmarkBookHandler_ListBooksHandler(b *testing.B) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type test struct {
		name     string
		args     args
		receiver *BookHandler
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
				bb.receiver.ListBooksHandler(
					bb.args.w,
					bb.args.r,
				)
			}
		})
	}
}
