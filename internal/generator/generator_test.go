package generator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
	"text/template"

	"github.com/kevindiu/gotest2/internal/models"
)

func TestGenerate(t *testing.T) {
	t.Parallel()
	type args struct {
		funcs              []*models.FunctionInfo
		imports            []string
		pkgName            string
		templatePath       string
		parallel           bool
		fuzz               bool
		benchmark          bool
		generateTests      bool
		existingTests      map[string]string
		entryPointTemplate string
	}
	type wants struct {
		want0   []byte
		wantErr error
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 []byte, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 []byte, gotErr error, tt *test) error {
		if gotErr != nil {
			return fmt.Errorf("unexpected error: %v", gotErr)
		}
		if !bytes.Contains(got0, tt.want.want0) {
			return fmt.Errorf("Generate() result = %s, want substring %s", got0, tt.want.want0)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "simple generation",
			init: func(t *testing.T, tt *test) {
				tmpDir := t.TempDir()
				tmplPath := filepath.Join(tmpDir, "test.tmpl")
				// Simple template that prints package name
				if err := os.WriteFile(tmplPath, []byte("package {{.PackageName}}\nfunc TestDummy() {}"), 0644); err != nil {
					t.Fatal(err)
				}
				tt.args.templatePath = tmplPath
				tt.args.pkgName = "mypkg"
				tt.args.funcs = []*models.FunctionInfo{{Name: "Foo"}}
				tt.want.want0 = []byte("package mypkg")
			},
		},
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
			got0, err := Generate(
				tt.args.funcs,
				tt.args.imports,
				tt.args.pkgName,
				tt.args.templatePath,
				tt.args.parallel,
				tt.args.fuzz,
				tt.args.benchmark,
				tt.args.generateTests,
				tt.args.existingTests,
				tt.args.entryPointTemplate,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, err, &tt); err != nil {
				t.Errorf("Generate() validation failed: %v", err)
			}
		})
	}
}

func Test_prepareMethods(t *testing.T) {
	t.Parallel()
	type args struct {
		funcs         []*models.FunctionInfo
		parallel      bool
		fuzz          bool
		benchmark     bool
		generateTests bool
		existingTests map[string]string
	}
	type wants struct {
		want0 []MethodData
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 []MethodData, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 []MethodData, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("prepareMethods() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "basic",
			args: args{
				funcs:         []*models.FunctionInfo{{Name: "Foo"}},
				generateTests: true,
			},
			validate: func(t *testing.T, got0 []MethodData, tt *test) error {
				if len(got0) != 1 {
					return fmt.Errorf("len = %d", len(got0))
				}
				if got0[0].Name != "Foo" {
					return fmt.Errorf("name = %s", got0[0].Name)
				}
				if !got0[0].Render {
					return fmt.Errorf("render = %v", got0[0].Render)
				}
				if !got0[0].HasLocalStruct {
					return fmt.Errorf("hasLocalStruct = %v", got0[0].HasLocalStruct)
				}
				return nil
			},
		},
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
			got0 := prepareMethods(
				tt.args.funcs,
				tt.args.parallel,
				tt.args.fuzz,
				tt.args.benchmark,
				tt.args.generateTests,
				tt.args.existingTests,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("prepareMethods() validation failed: %v", err)
			}
		})
	}
}

func Test_getTestFuncName(t *testing.T) {
	t.Parallel()
	type args struct {
		fn *models.FunctionInfo
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
			return fmt.Errorf("getTestFuncName() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "function",
			args: args{fn: &models.FunctionInfo{Name: "Foo", IsExported: true}},
			want: wants{want0: "Foo"},
		},
		{
			name: "method",
			args: args{fn: &models.FunctionInfo{
				Name: "Bar",
				Receiver: &models.Receiver{
					Type: "MyType",
				},
			}},
			want: wants{want0: "MyType_Bar"},
		},
		{
			name: "pointer receiver",
			args: args{fn: &models.FunctionInfo{
				Name: "Bar",
				Receiver: &models.Receiver{
					Type: "*MyType",
				},
			}},
			want: wants{want0: "MyType_Bar"},
		},
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
			got0 := getTestFuncName(
				tt.args.fn,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, &tt); err != nil {
				t.Errorf("getTestFuncName() validation failed: %v", err)
			}
		})
	}
}

func Test_loadTemplate(t *testing.T) {
	t.Parallel()
	type args struct {
		templatePath string
	}
	type wants struct {
		want0   *template.Template
		wantErr error
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 *template.Template, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 *template.Template, gotErr error, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("loadTemplate() got0 = %v, want %v", got0, tt.want.want0)
		}
		if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
			return fmt.Errorf("loadTemplate() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "valid template",
			init: func(t *testing.T, tt *test) {
				tmpDir := t.TempDir()
				tmplPath := filepath.Join(tmpDir, "test.tmpl")
				if err := os.WriteFile(tmplPath, []byte("Test"), 0644); err != nil {
					t.Fatal(err)
				}
				tt.args.templatePath = tmplPath
			},
			validate: func(t *testing.T, got0 *template.Template, gotErr error, tt *test) error {
				if gotErr != nil {
					return fmt.Errorf("unexpected error: %v", gotErr)
				}
				if got0 == nil {
					return fmt.Errorf("expected template, got nil")
				}
				return nil
			},
		},
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
			got0, err := loadTemplate(
				tt.args.templatePath,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, err, &tt); err != nil {
				t.Errorf("loadTemplate() validation failed: %v", err)
			}
		})
	}
}

func Test_executeTemplate(t *testing.T) {
	t.Parallel()
	type args struct {
		tmpl               *template.Template
		pkgName            string
		sourceImports      []string
		funcs              []MethodData
		templatePath       string
		entryPointTemplate string
	}
	type wants struct {
		want0   *bytes.Buffer
		wantErr error
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 *bytes.Buffer, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 *bytes.Buffer, gotErr error, tt *test) error {
		if gotErr != nil {
			return fmt.Errorf("unexpected error: %v", gotErr)
		}
		if got0.String() != tt.want.want0.String() {
			return fmt.Errorf("executeTemplate() got0 = %v, want %v", got0, tt.want.want0)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "basic execute",
			init: func(t *testing.T, tt *test) {
				tt.args.pkgName = "mypkg"
				tt.args.funcs = []MethodData{{FunctionInfo: &models.FunctionInfo{Name: "Foo"}}}
				// We need a template to execute. Load a simple one.
				tmpDir := t.TempDir()
				tmplPath := filepath.Join(tmpDir, "test.tmpl")
				if err := os.WriteFile(tmplPath, []byte("pkg: {{.PackageName}} func: {{range .Funcs}}{{.Name}}{{end}}"), 0644); err != nil {
					t.Fatal(err)
				}
				var err error
				tt.args.tmpl, err = loadTemplate(tmplPath)
				if err != nil {
					t.Fatal(err)
				}
				tt.args.templatePath = tmplPath
				tt.want.want0 = bytes.NewBufferString("pkg: mypkg func: Foo")
			},
		},
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
			got0, err := executeTemplate(
				tt.args.tmpl,
				tt.args.pkgName,
				tt.args.sourceImports,
				tt.args.funcs,
				tt.args.templatePath,
				tt.args.entryPointTemplate,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, err, &tt); err != nil {
				t.Errorf("executeTemplate() validation failed: %v", err)
			}
		})
	}
}

func Test_formatSource(t *testing.T) {
	t.Parallel()
	type args struct {
		buf *bytes.Buffer
	}
	type wants struct {
		want0   []byte
		wantErr error
	}
	type test struct {
		name     string
		args     args
		want     wants
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, got0 []byte, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, got0 []byte, gotErr error, tt *test) error {
		if !reflect.DeepEqual(got0, tt.want.want0) {
			return fmt.Errorf("formatSource() got0 = %s, want %s", got0, tt.want.want0)
		}
		if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
			return fmt.Errorf("formatSource() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "valid go code",
			args: args{buf: bytes.NewBufferString("package main\nfunc main() {}")},
			want: wants{want0: []byte("package main\n\nfunc main() {}\n")},
		},
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
			got0, err := formatSource(
				tt.args.buf,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, got0, err, &tt); err != nil {
				t.Errorf("formatSource() validation failed: %v", err)
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
		init     func(t *testing.T, tt *test)
		cleanup  func(t *testing.T, tt *test)
		validate func(t *testing.T, gotErr error, tt *test) error
	}
	defaultValidate := func(t *testing.T, gotErr error, tt *test) error {
		if fmt.Sprint(gotErr) != fmt.Sprint(tt.want.wantErr) {
			return fmt.Errorf("WriteFile() error = %v, wantErr %v", gotErr, tt.want.wantErr)
		}
		return nil
	}
	defaultInit := func(t *testing.T, tt *test) {}
	defaultCleanup := func(t *testing.T, tt *test) {}
	tests := []test{
		{
			name: "write new file",
			init: func(t *testing.T, tt *test) {
				tt.args.sourcePath = filepath.Join(t.TempDir(), "output.go")
				tt.args.content = []byte("package main")
			},
			validate: func(t *testing.T, gotErr error, tt *test) error {
				if gotErr != nil {
					return gotErr
				}
				// WriteFile appends _test.go
				base := strings.TrimSuffix(tt.args.sourcePath, ".go")
				testPath := base + "_test.go"
				if _, err := os.Stat(testPath); os.IsNotExist(err) {
					return fmt.Errorf("file %s not created", testPath)
				}
				return nil
			},
		},
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
			err := WriteFile(
				tt.args.sourcePath,
				tt.args.content,
			)
			if tt.validate == nil {
				tt.validate = defaultValidate
			}
			if err := tt.validate(t, err, &tt); err != nil {
				t.Errorf("WriteFile() validation failed: %v", err)
			}
		})
	}
}
