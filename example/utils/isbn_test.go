package utils

import (
	"fmt"
	"reflect"
	"testing"
)

func TestParseISBN(t *testing.T) {
	t.Parallel()
	type args struct {
		isbn string
	}
	type wants struct {
		want0 bool
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 bool, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 bool, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("ParseISBN() got0 = %v, want %v", got0, tt.want.want0)
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
			got0 := ParseISBN(
				tt.args.isbn,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("ParseISBN() validation failed: %v", err)
			}
		})
	}
}

func TestFormatISBN(t *testing.T) {
	t.Parallel()
	type args struct {
		raw string
	}
	type wants struct {
		want0 string
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 string, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 string, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("FormatISBN() got0 = %v, want %v", got0, tt.want.want0)
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
			got0 := FormatISBN(
				tt.args.raw,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("FormatISBN() validation failed: %v", err)
			}
		})
	}
}
