package handler

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"

	"github.com/kevindiu/gotest2/example/service"
)

func TestBookHandler_CreateBookHandler(t *testing.T) {
	t.Parallel()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type test struct {
		name     string
		receiver *BookHandler
		args     args
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, tt *test) error
	}
	defaultValidate := func(t *testing.T, tt *test) error {
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
			tt.receiver.CreateBookHandler(
				tt.args.w,
				tt.args.r,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, &tt); err != nil {
				t.Errorf("BookHandler_CreateBookHandler() validation failed: %v", err)
			}
		})
	}
}

func TestBookHandler_GetBookHandler(t *testing.T) {
	t.Parallel()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type test struct {
		name     string
		receiver *BookHandler
		args     args
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, tt *test) error
	}
	defaultValidate := func(t *testing.T, tt *test) error {
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
			tt.receiver.GetBookHandler(
				tt.args.w,
				tt.args.r,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, &tt); err != nil {
				t.Errorf("BookHandler_GetBookHandler() validation failed: %v", err)
			}
		})
	}
}

func TestBookHandler_ListBooksHandler(t *testing.T) {
	t.Parallel()
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	type test struct {
		name     string
		receiver *BookHandler
		args     args
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, tt *test) error
	}
	defaultValidate := func(t *testing.T, tt *test) error {
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
			tt.receiver.ListBooksHandler(
				tt.args.w,
				tt.args.r,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, &tt); err != nil {
				t.Errorf("BookHandler_ListBooksHandler() validation failed: %v", err)
			}
		})
	}
}

func TestNewBookHandler(t *testing.T) {
	t.Parallel()
	type args struct {
		svc *service.BookService
	}
	type wants struct {
		want0 *BookHandler
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 *BookHandler, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 *BookHandler, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("NewBookHandler() got0 = %v, want %v", got0, tt.want.want0)
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
			got0 := NewBookHandler(
				tt.args.svc,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("NewBookHandler() validation failed: %v", err)
			}
		})
	}
}
