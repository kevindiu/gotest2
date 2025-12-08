package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestChanSquare(t *testing.T) {
	t.Parallel()
	type args struct {
		in <-chan int
	}
	type wants struct {
		want0 <-chan int
	}
	type test struct {
		name     string
		args     args
		want     wants
		Init     func(t *testing.T, tt *test)
		Cleanup  func(t *testing.T, tt *test)
		Validate func(t *testing.T, got0 <-chan int, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 <-chan int, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("ChanSquare() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	tests := []test{
		{
			name: "Square 2, 3",
			args: args{in: make(chan int, 2)}, // Use buffered channel for correct Init population
			Init: func(t *testing.T, tt *test) {
				ch := make(chan int, 2)
				tt.args.in = ch
				ch <- 2
				ch <- 3
				close(ch)
			},
			Validate: func(t *testing.T, got0 <-chan int, tt *test) error {
				// Read from output channel
				res1, ok1 := <-got0
				if !ok1 || res1 != 4 {
					return fmt.Errorf("expected 4, got %d (ok=%v)", res1, ok1)
				}
				res2, ok2 := <-got0
				if !ok2 || res2 != 9 {
					return fmt.Errorf("expected 9, got %d (ok=%v)", res2, ok2)
				}
				_, ok3 := <-got0
				if ok3 {
					return fmt.Errorf("expected closed channel")
				}
				return nil
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
			got0 := ChanSquare(
				tt.args.in,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, &tt); err != nil {
				t.Errorf("ChanSquare() validation failed: %v", err)
			}
		})
	}
}

func TestDeferredExecution(t *testing.T) {
	t.Parallel()
	type args struct {
	}
	type wants struct {
		want0 string
	}
	type test struct {
		name     string
		args     args
		want     wants
		Init     func(t *testing.T, tt *test)
		Cleanup  func(t *testing.T, tt *test)
		Validate func(t *testing.T, got0 string, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 string, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("DeferredExecution() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	tests := []test{
		{
			name: "Check defer side effect",
			want: wants{want0: "hello world"},
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
			got0 := DeferredExecution()
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, &tt); err != nil {
				t.Errorf("DeferredExecution() validation failed: %v", err)
			}
		})
	}
}

func TestFuncFactory(t *testing.T) {
	t.Parallel()
	type args struct {
		n int
	}
	type wants struct {
		want0 func(int) int
	}
	type test struct {
		name     string
		args     args
		want     wants
		Init     func(t *testing.T, tt *test)
		Cleanup  func(t *testing.T, tt *test)
		Validate func(t *testing.T, got0 func(int) int, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 func(int) int, tt *test) error {
		return nil
	}
	tests := []test{
		{
			name: "Add 5",
			args: args{n: 5},
			Validate: func(t *testing.T, got0 func(int) int, tt *test) error {
				if res := got0(10); res != 15 {
					return fmt.Errorf("expected 15, got %d", res)
				}
				return nil
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
			got0 := FuncFactory(
				tt.args.n,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, &tt); err != nil {
				t.Errorf("FuncFactory() validation failed: %v", err)
			}
		})
	}
}

func TestMapKeys(t *testing.T) {
	t.Parallel()
	type args struct {
		m map[string]int
	}
	type wants struct {
		want0 []string
	}
	type test struct {
		name     string
		args     args
		want     wants
		Init     func(t *testing.T, tt *test)
		Cleanup  func(t *testing.T, tt *test)
		Validate func(t *testing.T, got0 []string, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 []string, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("MapKeys() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	tests := []test{
		{
			name: "Keys of {a:1, b:2}",
			args: args{m: map[string]int{"a": 1, "b": 2}},
			Validate: func(t *testing.T, got0 []string, tt *test) error {
				if len(got0) != 2 {
					return fmt.Errorf("expected 2 keys, got %d", len(got0))
				}
				// Check membership since order is random
				m := make(map[string]bool)
				for _, k := range got0 {
					m[k] = true
				}
				if !m["a"] || !m["b"] {
					return fmt.Errorf("expected keys [a, b], got %v", got0)
				}
				return nil
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
			got0 := MapKeys(
				tt.args.m,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, &tt); err != nil {
				t.Errorf("MapKeys() validation failed: %v", err)
			}
		})
	}
}

func TestVariadicSum(t *testing.T) {
	t.Parallel()
	type args struct {
		nums []int
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
			return fmt.Errorf("VariadicSum() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	tests := []test{
		{
			name: "Sum 1, 2, 3",
			args: args{nums: []int{1, 2, 3}},
			want: wants{want0: 6},
		},
		{
			name: "Sum empty",
			args: args{nums: []int{}},
			want: wants{want0: 0},
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
			got0 := VariadicSum(
				tt.args.nums...,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, &tt); err != nil {
				t.Errorf("VariadicSum() validation failed: %v", err)
			}
		})
	}
}
