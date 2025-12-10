package parser

import (
	"fmt"
	"go/ast"
	goparser "go/parser"
	"go/token"
	"go/types"
	"os"
	"path/filepath"
	"strings"

	"github.com/kevindiu/gotest2/internal/models"
	"golang.org/x/tools/go/packages"
)

// Parse parses the given patterns (file paths, directories, or ./... patterns)
// and returns a list of FileResult, organizing functions by their defining file.
func Parse(patterns []string) ([]*models.FileResult, error) {
	cfg := &packages.Config{
		Mode:  packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedSyntax | packages.NeedImports | packages.NeedFiles,
		Tests: false,
	}

	var pkgPatterns []string
	filePatternsByDir := make(map[string][]string)

	for _, p := range patterns {
		// Simple heuristic: if it ends in .go, treat as file.
		// Otherwise treat as package pattern.
		if strings.HasSuffix(p, ".go") {
			abs, err := filepath.Abs(p)
			if err != nil {
				return nil, fmt.Errorf("failed to get abs path for %s: %v", p, err)
			}
			dir := filepath.Dir(abs)
			filePatternsByDir[dir] = append(filePatternsByDir[dir], p)
		} else {
			pkgPatterns = append(pkgPatterns, p)
		}
	}

	var pkgs []*packages.Package

	// 1. Load generic package patterns (directories, ./...)
	if len(pkgPatterns) > 0 {
		p, err := packages.Load(cfg, pkgPatterns...)
		if err != nil {
			return nil, fmt.Errorf("failed to load packages: %v", err)
		}
		pkgs = append(pkgs, p...)
	}

	// 2. Load file patterns, grouped by directory
	for _, files := range filePatternsByDir {
		p, err := packages.Load(cfg, files...)
		if err != nil {
			return nil, fmt.Errorf("failed to load file packages: %v", err)
		}
		pkgs = append(pkgs, p...)
	}

	if packages.PrintErrors(pkgs) > 0 {

		return nil, fmt.Errorf("package load contained errors")
	}

	// Map to aggregate keys: absolute file path -> FileResult
	resultMap := make(map[string]*models.FileResult)

	for _, pkg := range pkgs {
		for _, file := range pkg.Syntax {
			pos := pkg.Fset.Position(file.Pos())
			// Skip test files
			if strings.HasSuffix(pos.Filename, "_test.go") {
				continue
			}

			// Capture imports for this file
			// We can't easily get *exact* used imports per function from AST without more work,
			// but the previous implementation accumulated imports per function via `qualifier`.
			// The `processFunction` uses `types.Func`, so we just need to get that from AST.

			for _, decl := range file.Decls {
				funcDecl, ok := decl.(*ast.FuncDecl)
				if !ok {
					continue
				}

				// Look up the type object
				obj := pkg.TypesInfo.Defs[funcDecl.Name]
				funcObj, ok := obj.(*types.Func)
				if !ok {
					continue
				}

				// Helper to add matching function to the map
				addFunc := func(path string, fn *models.FunctionInfo, importsMap map[string]struct{}) {
					if _, ok := resultMap[path]; !ok {
						resultMap[path] = &models.FileResult{
							Path:        path,
							PackageName: pkg.Name,
							Functions:   []*models.FunctionInfo{},
						}
					}
					resultMap[path].Functions = append(resultMap[path].Functions, fn)
					// Merge imports
					for imp := range importsMap {
						// Avoid duplicates
						found := false
						for _, existing := range resultMap[path].Imports {
							if existing == imp {
								found = true
								break
							}
						}
						if !found {
							resultMap[path].Imports = append(resultMap[path].Imports, imp)
						}
					}
				}

				funcs := []*models.FunctionInfo{}
				importsMap := make(map[string]struct{})
				processFunction(funcObj, &funcs, importsMap)
				if len(funcs) > 0 {
					addFunc(pos.Filename, funcs[0], importsMap)
				}
			}
		}
	}

	// Convert map to slice
	var results []*models.FileResult
	for _, res := range resultMap {
		results = append(results, res)
	}

	return results, nil
}

func processFunction(funcObj *types.Func, funcs *[]*models.FunctionInfo, importsMap map[string]struct{}) {
	sig, ok := funcObj.Type().(*types.Signature)
	if !ok {
		return
	}

	info := &models.FunctionInfo{
		Name:       funcObj.Name(),
		IsExported: funcObj.Exported(),
	}

	// Custom qualifier to ensure we get short package names
	qualifier := func(other *types.Package) string {
		if other == funcObj.Pkg() {
			return ""
		}
		// Capture import
		importsMap[other.Path()] = struct{}{}

		name := other.Name()
		if idx := strings.LastIndex(name, "/"); idx != -1 {
			name = name[idx+1:]
		}
		return name
	}

	extractTypeParams(sig, info, qualifier)
	extractReceiver(sig, info, qualifier)
	extractParams(sig, info, qualifier)
	extractResults(sig, info, qualifier)

	*funcs = append(*funcs, info)
}

func extractTypeParams(sig *types.Signature, info *models.FunctionInfo, qualifier types.Qualifier) {
	if tparams := sig.TypeParams(); tparams != nil && tparams.Len() > 0 {
		for i := 0; i < tparams.Len(); i++ {
			tp := tparams.At(i)
			info.TypeParams = append(info.TypeParams, &models.Field{
				Name:  tp.Obj().Name(),
				Type:  types.TypeString(tp.Constraint(), qualifier),
				Index: i,
			})
		}
	} else if recv := sig.Recv(); recv != nil {
		// Check if receiver has type params (e.g. method on generic type)
		var named *types.Named
		t := recv.Type()
		if ptr, ok := t.(*types.Pointer); ok {
			t = ptr.Elem()
		}
		if n, ok := t.(*types.Named); ok {
			named = n
		}

		if named != nil {
			if tparams := named.TypeParams(); tparams != nil && tparams.Len() > 0 {
				for i := 0; i < tparams.Len(); i++ {
					tp := tparams.At(i)
					info.TypeParams = append(info.TypeParams, &models.Field{
						Name:  tp.Obj().Name(),
						Type:  types.TypeString(tp.Constraint(), qualifier),
						Index: i,
					})
				}
			}
		}
	}
}

func extractReceiver(sig *types.Signature, info *models.FunctionInfo, qualifier types.Qualifier) {
	if recv := sig.Recv(); recv != nil {
		info.Receiver = &models.Receiver{
			Name: recv.Name(),
			Type: types.TypeString(recv.Type(), qualifier),
		}
	}
}

func extractParams(sig *types.Signature, info *models.FunctionInfo, qualifier types.Qualifier) {
	params := sig.Params()
	for i := 0; i < params.Len(); i++ {
		p := params.At(i)
		isVariadic := sig.Variadic() && i == params.Len()-1
		info.Parameters = append(info.Parameters, &models.Field{
			Name:       p.Name(),
			Type:       types.TypeString(p.Type(), qualifier),
			Index:      i,
			IsVariadic: isVariadic,
		})
	}
}

func extractResults(sig *types.Signature, info *models.FunctionInfo, qualifier types.Qualifier) {
	results := sig.Results()
	for i := 0; i < results.Len(); i++ {
		r := results.At(i)
		info.Results = append(info.Results, &models.Field{
			Name:  r.Name(),
			Type:  types.TypeString(r.Type(), qualifier),
			Index: i,
		})
	}
}

// ParseTests parses a test file and returns a map of test function names to their source code.
func ParseTests(path string) (map[string]string, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return nil, nil // proper error handling or just nil if no file
	}

	fset := token.NewFileSet()
	node, err := goparser.ParseFile(fset, path, nil, goparser.ParseComments)
	if err != nil {
		return nil, err
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	tests := make(map[string]string)
	for _, decl := range node.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		// We only care about TestXXX functions (public entry points)
		// We do NOT want to preserve runTestXXX or testXXXTestCase structs as those are generated/managed.
		if strings.HasPrefix(fn.Name.Name, "Test") {
			// Extract source
			start := fset.Position(fn.Pos()).Offset
			end := fset.Position(fn.End()).Offset
			tests[fn.Name.Name] = string(content[start:end])
		}
	}
	return tests, nil
}
