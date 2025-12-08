package generator

import (
	"reflect"
	"testing"

	"github.com/kevindiu/gotest2/internal/models"
)

func TestGenerate(t *testing.T) {
	t.Parallel()
	type args struct {
		funcs        []*models.FunctionInfo
		pkgName      string
		templatePath string
		parallel     bool
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			got0, err := Generate(
				tt.args.funcs,
				tt.args.pkgName,
				tt.args.templatePath,
				tt.args.parallel,
				false,
				nil,
			)

			if (err != nil) != tt.wantErr {
				t.Errorf("Generate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got0, tt.want) {
				t.Errorf("Generate() got0 = %v, want %v", got0, tt.want)
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
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			if err := WriteFile(
				tt.args.sourcePath,
				tt.args.content,
			); (err != nil) != tt.wantErr {
				t.Errorf("WriteFile() error = %v, wantErr %v", err, tt.wantErr)
			}

		})
	}
}
