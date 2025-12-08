package main

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMultiError(t *testing.T) {
	t.Parallel()
	type args struct {
		n int
	}
	type wants struct {
		want0    int
		wantErr1 error
		wantErr2 error
	}
	type test struct {
		name     string
		args     args
		want     wants
		Init     func(t *testing.T, tt *test)
		Cleanup  func(t *testing.T, tt *test)
		Validate func(t *testing.T, got0 int, gotErr1 error, gotErr2 error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 int, gotErr1 error, gotErr2 error, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("MultiError() got0 = %v, want %v", got0, tt.want.want0)
		}
		if fmt.Sprint(gotErr1) != fmt.Sprint(tt.want.wantErr1) {
			return fmt.Errorf("MultiError() error1 = %v, wantErr1 %v", gotErr1, tt.want.wantErr1)
		}
		if fmt.Sprint(gotErr2) != fmt.Sprint(tt.want.wantErr2) {
			return fmt.Errorf("MultiError() error2 = %v, wantErr2 %v", gotErr2, tt.want.wantErr2)
		}
		return nil
	}
	tests := []test{
		{
			name: "Negative input",
			args: args{n: -1},
			want: wants{
				want0:    0,
				wantErr1: fmt.Errorf("negative"),
			},
		},
		{
			name: "Too large input",
			args: args{n: 101},
			want: wants{
				want0:    0,
				wantErr2: fmt.Errorf("too large"),
			},
		},
		{
			name: "Normal input",
			args: args{n: 10},
			want: wants{
				want0: 10,
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
			got0, err1, err2 := MultiError(
				tt.args.n,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, err1, err2, &tt); err != nil {
				t.Errorf("MultiError() validation failed: %v", err)
			}
		})
	}
}
