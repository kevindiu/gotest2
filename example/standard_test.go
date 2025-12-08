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
		Init     func(t *testing.T, tt *test)
		Cleanup  func(t *testing.T, tt *test)
		Validate func(t *testing.T, got0 int, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 int, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("Add() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	tests := []test{
		{
			name: "1+1=2",
			args: args{a: 1, b: 1},
			want: wants{want0: 2},
		},
		{
			name: "10+20=30",
			args: args{a: 10, b: 20},
			want: wants{want0: 30},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.Init != nil {
				tt.Init(t, &tt)
			}
			if tt.Cleanup != nil {
				defer tt.Cleanup(t, &tt)
			}
			got0 := Add(
				tt.args.a,
				tt.args.b,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, &tt); err != nil {
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
		Init     func(t *testing.T, tt *test)
		Cleanup  func(t *testing.T, tt *test)
		Validate func(t *testing.T, got0 int, got1 int, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 int, got1 int, gotErr error, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("DivMod() got0 = %v, want %v", got0, tt.want.want0)
		}
		if !reflect.DeepEqual(got1, tt.want.want1) {
			return fmt.Errorf("DivMod() got1 = %v, want %v", got1, tt.want.want1)
		}
		if (gotErr != nil) != (tt.want.wantErr != nil) {
			return fmt.Errorf("DivMod() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		if gotErr != nil && tt.want.wantErr != nil && gotErr.Error() != tt.want.wantErr.Error() {
			return fmt.Errorf("DivMod() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		return nil
	}
	tests := []test{
		{
			name: "10/3",
			args: args{a: 10, b: 3},
			want: wants{want0: 3, want1: 1},
		},
		{
			name: "divide by zero",
			args: args{a: 10, b: 0},
			want: wants{
				want0:   0,
				want1:   0,
				wantErr: fmt.Errorf("division by zero"),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.Init != nil {
				tt.Init(t, &tt)
			}
			if tt.Cleanup != nil {
				defer tt.Cleanup(t, &tt)
			}
			got0, got1, err := DivMod(
				tt.args.a,
				tt.args.b,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, got1, err, &tt); err != nil {
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
		Init     func(t *testing.T, tt *test)
		Cleanup  func(t *testing.T, tt *test)
		Validate func(t *testing.T, got0 string, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 string, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("Person_Greet() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	tests := []test{
		{
			name: "Greet Alice",
			receiver: &Person{
				Name: "Alice",
				Age:  30,
			},
			args: args{
				msg: "Hello",
			},
			want: wants{
				want0: "Hello, I am Alice",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if tt.Init != nil {
				tt.Init(t, &tt)
			}
			if tt.Cleanup != nil {
				defer tt.Cleanup(t, &tt)
			}
			got0 := tt.receiver.Greet(
				tt.args.msg,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, &tt); err != nil {
				t.Errorf("Person_Greet() validation failed: %v", err)
			}
		})
	}
}
