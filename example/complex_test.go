package main

import (
	"fmt"
	"testing"
)

func TestParseConfig(t *testing.T) {
	t.Parallel()
	type args struct {
		data []byte
	}
	type wants struct {
		wantErr error
	}
	type test struct {
		name     string
		args     args
		want     wants
		Init     func(t *testing.T, tt *test)
		Cleanup  func(t *testing.T, tt *test)
		Validate func(t *testing.T, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, gotErr error, tt *test) error {
		if (gotErr != nil) != (tt.want.wantErr != nil) {
			return fmt.Errorf("ParseConfig() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		if gotErr != nil && tt.want.wantErr != nil && gotErr.Error() != tt.want.wantErr.Error() {
			return fmt.Errorf("ParseConfig() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		return nil
	}
	tests := []test{
		{
			name: "Valid JSON",
			args: args{data: []byte(`{"Name":"foo", "Val": 1}`)},
			want: wants{wantErr: nil},
		},
		{
			name: "Invalid JSON",
			args: args{data: []byte(`{invalid}`)},
			Validate: func(t *testing.T, gotErr error, tt *test) error {
				if gotErr == nil {
					return fmt.Errorf("expected error, got nil")
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
			err := ParseConfig(
				tt.args.data,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, err, &tt); err != nil {
				t.Errorf("ParseConfig() validation failed: %v", err)
			}
		})
	}
}
