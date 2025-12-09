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

// MethodData holds the data for a single test function generation/rendering.
type MethodData struct {
	*models.FunctionInfo
	Parallel       bool
	Render         bool   // If true, render the main TestXXX function.
	ExistingSource string // If Render is false, use this source code.
	Fuzz           bool   // If true, generate FuzzXXX target.
	Benchmark      bool   // If true, generate BenchmarkXXX target.
	HasLocalStruct bool   // If true, the test struct is defined locally (or will be).
}

// Generate generates test code for the given functions.
func Generate(funcs []*models.FunctionInfo, imports []string, pkgName string, templatePath string, parallel bool, fuzz bool, benchmark bool, generateTests bool, existingTests map[string]string, entryPointTemplate string) ([]byte, error) {
	if len(funcs) == 0 {
		return nil, fmt.Errorf("no functions to generate tests for")
	}

	wrappedFuncs := prepareMethods(funcs, parallel, fuzz, benchmark, generateTests, existingTests)

	tmpl, err := loadTemplate(templatePath)
	if err != nil {
		return nil, err
	}

	buf, err := executeTemplate(tmpl, pkgName, imports, wrappedFuncs, templatePath, entryPointTemplate)
	if err != nil {
		return nil, err
	}

	return formatSource(buf)
}

func prepareMethods(funcs []*models.FunctionInfo, parallel, fuzz, benchmark, generateTests bool, existingTests map[string]string) []MethodData {
	var wrappedFuncs []MethodData
	for _, fn := range funcs {
		testName := "Test" + getTestFuncName(fn)
		// Default to rendering if requested, unless it exists
		render := generateTests
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
			Benchmark:      benchmark,
			HasLocalStruct: hasLocalStruct,
		})
	}
	return wrappedFuncs
}

func getTestFuncName(fn *models.FunctionInfo) string {
	if fn.Receiver != nil {
		t := receiverName(fn.Receiver.Type)
		return fmt.Sprintf("%s_%s", t, fn.Name)
	}

	// For unexported functions without valid receiver, use _prefix
	// This results in Test_funcName
	if !fn.IsExported {
		return "_" + fn.Name
	}
	return fn.Name
}

func loadTemplate(templatePath string) (*template.Template, error) {
	funcMap := FuncMap()
	if templatePath != "" {
		name := filepath.Base(templatePath)
		tmpl, err := template.New(name).Funcs(funcMap).ParseFiles(templatePath)
		if err != nil {
			return nil, fmt.Errorf("failed to parse custom template %s: %v", templatePath, err)
		}
		return tmpl, nil
	}
	tmpl, err := template.New("body.tmpl").Funcs(funcMap).ParseFS(templates.FS, "*.tmpl")
	if err != nil {
		return nil, fmt.Errorf("failed to parse embedded templates: %v", err)
	}
	return tmpl, nil
}

func executeTemplate(tmpl *template.Template, pkgName string, sourceImports []string, funcs []MethodData, templatePath, entryPointTemplate string) (*bytes.Buffer, error) {
	// Merge default imports with source imports
	// Use map to deduplicate
	importsMap := make(map[string]struct{})
	defaults := []string{"testing", "reflect"}
	for _, imp := range defaults {
		importsMap[imp] = struct{}{}
	}
	for _, imp := range sourceImports {
		// Clean import paths (remove quotes if any, though model likely has raw paths)
		// Check if it's not empty
		if imp != "" {
			importsMap[imp] = struct{}{}
		}
	}

	var importsList []string
	for imp := range importsMap {
		importsList = append(importsList, imp)
	}

	data := struct {
		PackageName string
		Imports     []string
		Funcs       []MethodData
	}{
		PackageName: pkgName,
		Imports:     importsList,
		Funcs:       funcs,
	}

	var buf bytes.Buffer
	var err error
	if templatePath != "" {
		err = tmpl.Execute(&buf, data)
	} else {
		if entryPointTemplate == "" {
			entryPointTemplate = "body.tmpl"
		}
		err = tmpl.ExecuteTemplate(&buf, entryPointTemplate, data)
	}

	if err != nil {
		return nil, fmt.Errorf("failed to execute template: %v", err)
	}
	return &buf, nil
}

func formatSource(buf *bytes.Buffer) ([]byte, error) {
	// Format the code using imports.Process to fix imports
	src, err := imports.Process("", buf.Bytes(), nil)
	if err != nil {
		// Return unformatted code for debugging if format fails
		return buf.Bytes(), fmt.Errorf("failed to format source: %v", err)
	}

	// If the generated code doesn't contain any key indicators of test code (e.g. "func "),
	// assume nothing meaningful was generated (e.g. only header).
	if !bytes.Contains(src, []byte("func ")) {
		return nil, nil
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
