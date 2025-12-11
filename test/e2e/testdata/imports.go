package testdata

import (
	"fmt"
	"strings"
)

func Format(s string) string {
	return fmt.Sprintf("Prefix: %s", s)
}

func Upper(s string) string {
	return strings.ToUpper(s)
}
