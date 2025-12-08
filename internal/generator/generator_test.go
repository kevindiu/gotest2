package generator

import (
	"fmt"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/kevindiu/gotest2/internal/models"
)

func TestGenerate(t *testing.T) {
	t.Parallel()
	type args struct {
		funcs         []*models.FunctionInfo
		pkgName       string
		templatePath  string
		parallel      bool
		fuzz          bool
		existingTests map[string]string
	}
	type wants struct {
		want0   []byte
		wantErr error
	}
	type test struct {
		name     string
		args     args
		want     wants
		Init     func(t *testing.T, tt *test)
		Cleanup  func(t *testing.T, tt *test)
		Validate func(t *testing.T, got0 []byte, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 []byte, gotErr error, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("Generate() got0 = %v, want %v", got0, tt.want.want0)
		}
		if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
			return fmt.Errorf("Generate() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		return nil
	}
	tests := []test{
		{
			name: "Generate simple function",
			args: args{
				funcs: []*models.FunctionInfo{
					{
						Name:       "Add",
						IsExported: true,
						Parameters: []*models.Field{
							{Name: "a", Type: "int"},
							{Name: "b", Type: "int"},
						},
						Results: []*models.Field{
							{Name: "", Type: "int"},
						},
					},
				},
				pkgName:  "example",
				parallel: true,
			},
			Validate: func(t *testing.T, got0 []byte, gotErr error, tt *test) error {
				if gotErr != nil {
					return fmt.Errorf("unexpected error: %v", gotErr)
				}
				s := string(got0)
				if s == "" {
					return fmt.Errorf("expected generated code, got empty string")
				}
				expected := []string{"package example", "func TestAdd(t *testing.T)", "t.Parallel()"}
				for _, e := range expected {
					if !strings.Contains(s, e) {
						return fmt.Errorf("missing expected string: %s", e)
					}
				}
				return nil
			},
		},
		{
			name: "Generate method on struct",
			args: args{
				funcs: []*models.FunctionInfo{
					{
						Name:       "Greet",
						IsExported: true,
						Receiver: &models.Receiver{
							Name: "p",
							Type: "*Person",
						},
						Parameters: []*models.Field{
							{Name: "msg", Type: "string"},
						},
						Results: []*models.Field{
							{Name: "", Type: "string"},
						},
					},
				},
				pkgName:  "example",
				parallel: true,
			},
			Validate: func(t *testing.T, got0 []byte, gotErr error, tt *test) error {
				if gotErr != nil {
					return fmt.Errorf("unexpected error: %v", gotErr)
				}
				s := string(got0)
				// We expect the STRUCT to represent the receiver, but the test slice will be empty/TODO.
				expected := []string{
					"func TestPerson_Greet(t *testing.T)",
					"type test struct",
					"receiver *Person",
				}
				for _, e := range expected {
					if !strings.Contains(s, e) {
						return fmt.Errorf("missing expected string: '%s'", e)
					}
				}
				return nil
			},
		},
		{
			name: "Generate Generic Function",
			args: args{
				funcs: []*models.FunctionInfo{
					{
						Name:       "GenericSum",
						IsExported: true,
						TypeParams: []*models.Field{
							{Name: "T", Type: "int | float64"},
						},
						Parameters: []*models.Field{
							{Name: "a", Type: "T"},
							{Name: "b", Type: "T"},
						},
						Results: []*models.Field{
							{Name: "", Type: "T"},
						},
					},
				},
				pkgName:  "example",
				parallel: true,
			},
			Validate: func(t *testing.T, got0 []byte, gotErr error, tt *test) error {
				if gotErr != nil {
					return fmt.Errorf("unexpected error: %v", gotErr)
				}
				s := string(got0)
				expectedSignatures := []string{
					"func runTestGenericSum[T int | float64](t *testing.T, cases []testGenericSumTestCase[T])",
					"type testGenericSumTestCase[T int | float64] struct",
				}
				for _, e := range expectedSignatures {
					if !strings.Contains(s, e) {
						return fmt.Errorf("missing expected string: %s", e)
					}
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
			got0, err := Generate(
				tt.args.funcs,
				tt.args.pkgName,
				tt.args.templatePath,
				tt.args.parallel,
				tt.args.fuzz,
				tt.args.existingTests,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, got0, err, &tt); err != nil {
				t.Errorf("Generate() validation failed: %v", err)
			}
		})
	}
}

func TestWriteFile(t *testing.T) {
	t.Parallel()
	type args struct {
		sourcePath string
		content    []byte
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
		if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
			return fmt.Errorf("WriteFile() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		return nil
	}
	tests := []test{
		{
			name: "Write to temp file",
			args: args{
				sourcePath: "temp_test_gen.go",
				content:    []byte("package main\n"),
			},
			Validate: func(t *testing.T, gotErr error, tt *test) error {
				if gotErr != nil {
					return fmt.Errorf("unexpected error: %v", gotErr)
				}
				// Verify file exists
				if _, err := os.Stat("temp_test_gen_test.go"); os.IsNotExist(err) {
					return fmt.Errorf("expected file temp_test_gen_test.go to exist")
				}
				return nil
			},
			Cleanup: func(t *testing.T, tt *test) {
				os.Remove("temp_test_gen_test.go")
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
			err := WriteFile(
				tt.args.sourcePath,
				tt.args.content,
			)
			validation := defaultValidate
			if tt.Validate != nil {
				validation = tt.Validate
			}
			if err := validation(t, err, &tt); err != nil {
				t.Errorf("WriteFile() validation failed: %v", err)
			}
		})
	}
}
