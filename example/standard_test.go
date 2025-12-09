package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAdd(t *testing.T) {
	t.Parallel()
	type args struct {
		a int
		b int
	}
	type wants struct {
		want0 int
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 int, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 int, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("Add() got0 = %v, want %v", got0, tt.want.want0)
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
			got0 := Add(
				tt.args.a,
				tt.args.b,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("Add() validation failed: %v", err)
			}
		})
	}
}

func TestDivMod(t *testing.T) {
	t.Parallel()
	type args struct {
		a int
		b int
	}
	type wants struct {
		want0   int
		want1   int
		wantErr error
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 int, got1 int, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 int, got1 int, gotErr error, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("DivMod() got0 = %v, want %v", got0, tt.want.want0)
		}
		if !reflect.DeepEqual(got1, tt.want.want1) {
			return fmt.Errorf("DivMod() got1 = %v, want %v", got1, tt.want.want1)
		}
		if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
			return fmt.Errorf("DivMod() error = %v, wantErr %v", gotErr, tt.want.wantErr)
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
			got0, got1, err := DivMod(
				tt.args.a,
				tt.args.b,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, got1, err, &tt); err != nil {
				t.Errorf("DivMod() validation failed: %v", err)
			}
		})
	}
}

func TestPerson_Greet(t *testing.T) {
	t.Parallel()
	type args struct {
		msg string
	}
	type wants struct {
		want0 string
	}
	type test struct {
		name     string
		receiver *Person
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 string, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 string, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("Person_Greet() got0 = %v, want %v", got0, tt.want.want0)
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
			got0 := tt.receiver.Greet(
				tt.args.msg,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("Person_Greet() validation failed: %v", err)
			}
		})
	}
}
