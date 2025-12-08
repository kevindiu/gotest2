package parser

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/kevindiu/gotest2/internal/models"
)

func TestParse(t *testing.T) {
	t.Parallel()
	type args struct {
		patterns []string
	}
	type wants struct {
		want0   []*models.FileResult
		wantErr error
	}
	type test struct {
		name     string
		args     args
		want     wants
		Init     func(t *testing.T, tt *test)
		Cleanup  func(t *testing.T, tt *test)
		Validate func(t *testing.T, got0 []*models.FileResult, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 []*models.FileResult, gotErr error, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("Parse() got0 = %v, want %v", got0, tt.want.want0)
		}
		if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
			return fmt.Errorf("Parse() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		return nil
	}
	tests := []test{
		{
			name: "Valid file",
			args: args{
				patterns: []string{"../../example/standard.go"},
			},
			Validate: func(t *testing.T, got0 []*models.FileResult, gotErr error, tt *test) error {
				if gotErr != nil {
					return fmt.Errorf("unexpected error: %v", gotErr)
				}
				if len(got0) != 1 {
					return fmt.Errorf("expected 1 file result, got %d", len(got0))
				}
				if len(got0[0].Functions) == 0 {
					return fmt.Errorf("expected functions, got 0")
				}
				// Verify specific function exists
				found := false
				for _, fn := range got0[0].Functions {
					if fn.Name == "Add" {
						found = true
						break
					}
				}
				if !found {
					return fmt.Errorf("expected function Add not found")
				}
				return nil
			},
		},
		{
			name: "Non-existent file",
			args: args{
				patterns: []string{"nonexistent.go"},
			},
			Validate: func(t *testing.T, got0 []*models.FileResult, gotErr error, tt *test) error {
				if gotErr == nil {
					return fmt.Errorf("expected error for non-existent file, got nil")
				}
				return nil
			},
		},
		{
			name: "Parse Generics",
			args: args{
				patterns: []string{"../../example/generics.go"},
			},
			Validate: func(t *testing.T, got0 []*models.FileResult, gotErr error, tt *test) error {
				if gotErr != nil {
					return fmt.Errorf("unexpected error: %v", gotErr)
				}
				if len(got0) != 1 {
					return fmt.Errorf("expected 1 file result, got %d", len(got0))
				}
				// Verify GenericSum
				var genericSum *models.FunctionInfo
				for _, fn := range got0[0].Functions {
					if fn.Name == "GenericSum" {
						genericSum = fn
						break
					}
				}
				if genericSum == nil {
					return fmt.Errorf("GenericSum not found")
				}
				if len(genericSum.TypeParams) == 0 {
					return fmt.Errorf("expected GenericSum to have type params")
				}
				// Check generic method MyList.Add
				var myListAdd *models.FunctionInfo
				for _, fn := range got0[0].Functions {
					if fn.Name == "Add" && fn.Receiver.Name == "l" {
						myListAdd = fn
						break
					}
				}
				if myListAdd == nil {
					return fmt.Errorf("MyList.Add not found")
				}
				if !strings.Contains(myListAdd.Receiver.Type, "MyList[T]") {
					return fmt.Errorf("expected receiver type *MyList[T], got %s", myListAdd.Receiver.Type)
				}
				return nil
			},
		},
		{
			name: "Parse Standard Methods",
			args: args{
				patterns: []string{"../../example/standard.go"},
			},
			Validate: func(t *testing.T, got0 []*models.FileResult, gotErr error, tt *test) error {
				if gotErr != nil {
					return fmt.Errorf("unexpected error: %v", gotErr)
				}
				// Find Person.Greet
				var greet *models.FunctionInfo
				for _, f := range got0 {
					for _, fn := range f.Functions {
						if fn.Name == "Greet" {
							greet = fn
							break
						}
					}
				}
				if greet == nil {
					return fmt.Errorf("Greet method not found")
				}
				if greet.Receiver == nil {
					return fmt.Errorf("expected receiver for Greet")
				}
				if !strings.Contains(greet.Receiver.Type, "Person") {
					return fmt.Errorf("expected receiver type *Person, got %s", greet.Receiver.Type)
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
			got0, err := Parse(
				tt.args.patterns,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, err, &tt); err != nil {
				t.Errorf("Parse() validation failed: %v", err)
			}
		})
	}
}

func TestParseTests(t *testing.T) {
	t.Parallel()
	type args struct {
		path string
	}
	type wants struct {
		want0   map[string]string
		wantErr error
	}
	type test struct {
		name     string
		args     args
		want     wants
		Init     func(t *testing.T, tt *test)
		Cleanup  func(t *testing.T, tt *test)
		Validate func(t *testing.T, got0 map[string]string, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 map[string]string, gotErr error, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("ParseTests() got0 = %v, want %v", got0, tt.want.want0)
		}
		if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
			return fmt.Errorf("ParseTests() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		return nil
	}
	tests := []test{
		{
			name: "Valid test file",
			args: args{path: "../../example/standard_test.go"},
			Validate: func(t *testing.T, got0 map[string]string, gotErr error, tt *test) error {
				if gotErr != nil {
					return fmt.Errorf("unexpected error: %v", gotErr)
				}
				if _, ok := got0["TestAdd"]; !ok {
					return fmt.Errorf("expected TestAdd to be found")
				}
				return nil
			},
		},
		{
			name: "Non-existent file",
			args: args{path: "nonexistent_test.go"},
			want: wants{
				want0:   nil,
				wantErr: nil, // Current implementation returns nil map and nil error for non-exist? Check implementation.
			},
			Validate: func(t *testing.T, got0 map[string]string, gotErr error, tt *test) error {
				// Implementation returns nil, nil for non-existent file?
				// "if _, err := os.Stat(path); os.IsNotExist(err) { return nil, nil }"
				if gotErr != nil {
					return fmt.Errorf("unexpected error: %v", gotErr)
				}
				if got0 != nil {
					return fmt.Errorf("expected nil result, got %v", got0)
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
			got0, err := ParseTests(
				tt.args.path,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, err, &tt); err != nil {
				t.Errorf("ParseTests() validation failed: %v", err)
			}
		})
	}
}
