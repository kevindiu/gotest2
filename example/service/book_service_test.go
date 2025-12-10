package service

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/kevindiu/gotest2/example/model"
	"github.com/kevindiu/gotest2/example/repository"
)

func TestBookService_CreateBook(t *testing.T) {
	t.Parallel()
	type args struct {
		title  string
		author string
		isbn   string
	}
	type wants struct {
		want0   model.Book
		wantErr error
	}
	type test struct {
		name     string
		receiver *BookService
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 model.Book, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 model.Book, gotErr error, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("BookService_CreateBook() got0 = %v, want %v", got0, tt.want.want0)
		}
		if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
			return fmt.Errorf("BookService_CreateBook() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.init == nil {
				tt.init = defaultInit
			}
			tt.init(t, &tt)
			if tt.cleanup == nil {
				tt.cleanup = defaultCleanup
			}
			defer tt.cleanup(t, &tt)
			got0, err := tt.receiver.CreateBook(
				tt.args.title,
				tt.args.author,
				tt.args.isbn,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, err, &tt); err != nil {
				t.Errorf("BookService_CreateBook() validation failed: %v", err)
			}
		})
	}
}

func TestBookService_GetBook(t *testing.T) {
	t.Parallel()
	type args struct {
		id string
	}
	type wants struct {
		want0   model.Book
		wantErr error
	}
	type test struct {
		name     string
		receiver *BookService
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 model.Book, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 model.Book, gotErr error, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("BookService_GetBook() got0 = %v, want %v", got0, tt.want.want0)
		}
		if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
			return fmt.Errorf("BookService_GetBook() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.init == nil {
				tt.init = defaultInit
			}
			tt.init(t, &tt)
			if tt.cleanup == nil {
				tt.cleanup = defaultCleanup
			}
			defer tt.cleanup(t, &tt)
			got0, err := tt.receiver.GetBook(
				tt.args.id,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, err, &tt); err != nil {
				t.Errorf("BookService_GetBook() validation failed: %v", err)
			}
		})
	}
}

func TestBookService_ListBooks(t *testing.T) {
	t.Parallel()
	type args struct {
	}
	type wants struct {
		want0   []model.Book
		wantErr error
	}
	type test struct {
		name     string
		receiver *BookService
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 []model.Book, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 []model.Book, gotErr error, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("BookService_ListBooks() got0 = %v, want %v", got0, tt.want.want0)
		}
		if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
			return fmt.Errorf("BookService_ListBooks() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.init == nil {
				tt.init = defaultInit
			}
			tt.init(t, &tt)
			if tt.cleanup == nil {
				tt.cleanup = defaultCleanup
			}
			defer tt.cleanup(t, &tt)
			got0, err := tt.receiver.ListBooks()
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, err, &tt); err != nil {
				t.Errorf("BookService_ListBooks() validation failed: %v", err)
			}
		})
	}
}

func TestNewBookService(t *testing.T) {
	t.Parallel()
	type args struct {
		repo repository.Repository[model.Book, string]
	}
	type wants struct {
		want0 *BookService
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 *BookService, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 *BookService, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("NewBookService() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.init == nil {
				tt.init = defaultInit
			}
			tt.init(t, &tt)
			if tt.cleanup == nil {
				tt.cleanup = defaultCleanup
			}
			defer tt.cleanup(t, &tt)
			got0 := NewBookService(
				tt.args.repo,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("NewBookService() validation failed: %v", err)
			}
		})
	}
}
