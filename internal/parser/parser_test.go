package parser

import (
	"testing"
)

func TestParse(t *testing.T) {
	t.Parallel()
	type args struct {
		patterns []string
	}
	tests := []struct {
		name    string
		args    args
		wantLen int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			results, err := Parse(
				tt.args.patterns,
			)

			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if len(results) != tt.wantLen {
				t.Errorf("Parse() got %d results, want %d", len(results), tt.wantLen)
			}
		})
	}
}
