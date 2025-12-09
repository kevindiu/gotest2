package generator

import (
	"strings"
	"text/template"

	"github.com/kevindiu/gotest2/internal/models"
)

// FuncMap returns the func map for templates.
func FuncMap() template.FuncMap {
	return template.FuncMap{
		"add":          add,
		"receiverName": receiverName,
		"isFuzzable":   isFuzzable,
		"isFunc":       isFunc,
	}
}

func add(a, b int) int {
	return a + b
}

func receiverName(t string) string {
	// Type string might be "*MyList[int]" or "Person"
	cleanType := t
	cleanType = strings.TrimPrefix(cleanType, "*")
	if idx := strings.Index(cleanType, "["); idx != -1 {
		cleanType = cleanType[:idx]
	}
	// Remove package prefix if any (e.g. models.Person -> Person)
	if idx := strings.LastIndex(cleanType, "."); idx != -1 {
		cleanType = cleanType[idx+1:]
	}

	if len(cleanType) > 0 {
		cleanType = strings.ToUpper(cleanType[:1]) + cleanType[1:]
	}
	return cleanType
}

func isFuzzable(t string, typeParams []*models.Field) bool {
	// Check if type matches a Type Parameter
	for _, tp := range typeParams {
		if t == tp.Name {
			return true
		}
	}

	// Check if type is supported by testing.F
	// string, []byte, int, int8, int16, int32/rune, int64, uint, uint8/byte, uint16, uint32, uint64, float32, float64, bool
	switch t {
	case "string", "[]byte", "int", "int8", "int16", "int32", "rune", "int64",
		"uint", "uint8", "byte", "uint16", "uint32", "uint64",
		"float32", "float64", "bool":
		return true
	}
	return false
}

func isFunc(t string) bool {
	return strings.HasPrefix(t, "func")
}
