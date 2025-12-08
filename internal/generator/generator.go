package generator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"golang.org/x/tools/imports"

	"github.com/kevindiu/gotest2/internal/models"
	"github.com/kevindiu/gotest2/templates"
)

// Generate generates test code for the given functions.
func Generate(funcs []*models.FunctionInfo, pkgName string, templatePath string, parallel bool, fuzz bool, existingTests map[string]string) ([]byte, error) {
	if len(funcs) == 0 {
		return nil, fmt.Errorf("no functions to generate tests for")
	}

	// Determine package name (naive approach, should really use parsed package name)
	// For now assuming we are in the same package so package name is consistent
	// or we are generating for a "_test" package.
	// Let's use "package_test" pattern or just "package" depending on user pref.
	// I'll default to same package for now for simplicity or 'models' if not found?
	// Actually, I should pass the package name from parser.
	// Let's modify models to include PackageName.

	// For this step I'll assume we pass it or extract it.
	// Let's blindly use "main" or extract from sourcePath?
	// Recommendation: Update parser to return PackageName.
	// But for now, I'll allow the template to handle it or just hardcode/placeholder.

	// Wrap functions with generation context (Parallel flag)
	type MethodData struct {
		*models.FunctionInfo
		Parallel       bool
		Render         bool   // If true, render the main TestXXX function.
		ExistingSource string // If Render is false, use this source code.
		Fuzz           bool   // If true, generate FuzzXXX target.
		HasLocalStruct bool   // If true, the test struct is defined locally (or will be).
	}
	var wrappedFuncs []MethodData

	// Helper to calculate test function name to check against existing tests
	// This logic duplicates the template logic, which is a bit fragile but necessary for this approach.
	getTestFuncName := func(fn *models.FunctionInfo) string {
		name := fn.Name
		if fn.Receiver != nil {
			// Logic from template: receiverName .Receiver.Type
			// receiverName removes *, [], package.
			// Replicate receiverName logic:
			t := fn.Receiver.Type
			t = strings.TrimPrefix(t, "*")
			if idx := strings.Index(t, "["); idx != -1 {
				t = t[:idx]
			}
			if idx := strings.LastIndex(t, "."); idx != -1 {
				t = t[idx+1:]
			}
			if len(t) > 0 {
				t = strings.ToUpper(t[:1]) + t[1:]
			}
			return fmt.Sprintf("%s_%s", t, name)
		}
		return name
	}

	for _, fn := range funcs {
		testName := "Test" + getTestFuncName(fn)
		// Default to rendering unless it exists
		render := true
		existingSource := ""
		hasLocalStruct := true // Default to true for new tests (we generate local struct)

		if src, exists := existingTests[testName]; exists {
			render = false
			existingSource = src
			// Check if the existing test source contains the local struct definition
			structName := fmt.Sprintf("type test%sTestCase struct", getTestFuncName(fn))
			if strings.Contains(src, structName) || strings.Contains(src, "type test struct") {
				hasLocalStruct = true
			} else {
				hasLocalStruct = false
			}
		}

		wrappedFuncs = append(wrappedFuncs, MethodData{
			FunctionInfo:   fn,
			Parallel:       parallel,
			Render:         render,
			ExistingSource: existingSource,
			Fuzz:           fuzz,
			HasLocalStruct: hasLocalStruct,
		})
	}

	data := struct {
		PackageName string
		Imports     []string
		Funcs       []MethodData
	}{
		PackageName: pkgName, // Use the parsed package name (usually main or x)
		// If we want x_test, we should append _test?
		// But for verification let's just use pkgName and we can move file to pkgName_test later or user preference.
		// Actually, if we write to _test.go, we should probably use pkgName + "_test" or just pkgName if it's internal test.
		// Let's use pkgName currently.
		Imports: []string{"testing", "reflect"},
		Funcs:   wrappedFuncs,
	}

	// Load templates
	var tmpl *template.Template
	var err error

	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"receiverName": func(t string) string {
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
		},
		"isFuzzable": func(t string) bool {
			// Check if type is supported by testing.F
			// string, []byte, int, int8, int16, int32/rune, int64, uint, uint8/byte, uint16, uint32, uint64, float32, float64, bool
			switch t {
			case "string", "[]byte", "int", "int8", "int16", "int32", "rune", "int64",
				"uint", "uint8", "byte", "uint16", "uint32", "uint64",
				"float32", "float64", "bool":
				return true
			}
			return false
		},
		"isFunc": func(t string) bool {
			return strings.HasPrefix(t, "func")
		},
	}

	if templatePath != "" {
		// Custom template overrides everything for now (as per previous logic)
		// Or should we allow partial overrides? User asked for "custom template", usually implies full control.
		// For backward compatibility with the CLI flag I just added, if a custom template is provided,
		// we use it as the single source of truth.
		name := filepath.Base(templatePath)
		tmpl, err = template.New(name).Funcs(funcMap).ParseFiles(templatePath)
		if err != nil {
			return nil, fmt.Errorf("failed to parse custom template %s: %v", templatePath, err)
		}
	} else {
		// Use embedded templates
		tmpl, err = template.New("body.tmpl").Funcs(funcMap).ParseFS(templates.FS, "*.tmpl")
		if err != nil {
			return nil, fmt.Errorf("failed to parse embedded templates: %v", err)
		}
	}

	var buf bytes.Buffer
	// Execute the "body.tmpl" template which is the entry point
	// If custom template is used, we execute that one.
	// Note: If custom template is used, the name depends on filepath.Base(templatePath).
	// We should probably just Execute whatever is the main one.

	if templatePath != "" {
		err = tmpl.Execute(&buf, data)
	} else {
		err = tmpl.ExecuteTemplate(&buf, "body.tmpl", data)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %v", err)
	}

	// Format the code using imports.Process to fix imports
	src, err := imports.Process("", buf.Bytes(), nil)
	if err != nil {
		// Return unformatted code for debugging if format fails
		return buf.Bytes(), fmt.Errorf("failed to format source: %v", err)
	}

	return src, nil
}

// WriteFile writes the generated code to the test file.
func WriteFile(sourcePath string, content []byte) error {
	ext := filepath.Ext(sourcePath)
	base := sourcePath[:len(sourcePath)-len(ext)]
	testPath := base + "_test.go"

	return os.WriteFile(testPath, content, 0644)
}
