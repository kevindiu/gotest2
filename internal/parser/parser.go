package parser

import (
	"fmt"
	"go/ast"
	goparser "go/parser"
	"go/token"
	"go/types"
	"os"
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

	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		return nil, fmt.Errorf("failed to load packages: %v", err)
	}
	if packages.PrintErrors(pkgs) > 0 {
		return nil, fmt.Errorf("package load contained errors")
	}

	// Map to aggregate keys: absolute file path -> FileResult
	resultMap := make(map[string]*models.FileResult)

	for _, pkg := range pkgs {
		scope := pkg.Types.Scope()
		for _, name := range scope.Names() {
			obj := scope.Lookup(name)

			// Helper to add matching function to the map
			addFunc := func(path string, fn *models.FunctionInfo) {
				if _, ok := resultMap[path]; !ok {
					resultMap[path] = &models.FileResult{
						Path:        path,
						PackageName: pkg.Name,
						Functions:   []*models.FunctionInfo{},
					}
				}
				resultMap[path].Functions = append(resultMap[path].Functions, fn)
			}

			if funcObj, ok := obj.(*types.Func); ok {
				// Top-level function
				pos := pkg.Fset.Position(funcObj.Pos())
				// Only include if it's in the loaded package (not imported)
				// and not in a test file (sanity check, usually filtered by Tests: false logic but safe to keep)
				if !strings.HasSuffix(pos.Filename, "_test.go") {
					funcs := []*models.FunctionInfo{}
					processFunction(funcObj, &funcs)
					if len(funcs) > 0 {
						addFunc(pos.Filename, funcs[0])
					}
				}
			}

			// Methods on Named types
			if typeNameObj, ok := obj.(*types.TypeName); ok {
				if named, ok := typeNameObj.Type().(*types.Named); ok {
					for i := 0; i < named.NumMethods(); i++ {
						method := named.Method(i)
						pos := pkg.Fset.Position(method.Pos())
						if !strings.HasSuffix(pos.Filename, "_test.go") {
							funcs := []*models.FunctionInfo{}
							processFunction(method, &funcs)
							if len(funcs) > 0 {
								addFunc(pos.Filename, funcs[0])
							}
						}
					}
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

func processFunction(funcObj *types.Func, funcs *[]*models.FunctionInfo) {
	sig, ok := funcObj.Type().(*types.Signature)
	if !ok {
		return
	}

	info := &models.FunctionInfo{
		Name:       funcObj.Name(),
		IsExported: funcObj.Exported(),
	}

	// Type Constraints (Generics)
	if tparams := sig.TypeParams(); tparams != nil && tparams.Len() > 0 {
		for i := 0; i < tparams.Len(); i++ {
			tp := tparams.At(i)
			// simplified type representation
			info.TypeParams = append(info.TypeParams, &models.Field{
				Name:  tp.Obj().Name(),
				Type:  tp.Constraint().String(),
				Index: i,
			})
		}
	} else if recv := sig.Recv(); recv != nil {
		// Check if receiver has type params (e.g. method on generic type)
		// Receiver type might be *Named or Named
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
						Type:  tp.Constraint().String(),
						Index: i,
					})
				}
			}
		}
	}

	// Custom qualifier to ensure we get short package names (e.g. "models" instead of "github.com/.../models")
	// This allows goimports to resolve the import path later.
	qualifier := func(other *types.Package) string {
		if other == funcObj.Pkg() {
			return ""
		}
		// Paranoid check: ensure name doesn't contain slashes
		name := other.Name()
		if idx := strings.LastIndex(name, "/"); idx != -1 {
			name = name[idx+1:]
		}
		return name
	}

	// Receiver
	if recv := sig.Recv(); recv != nil {
		info.Receiver = &models.Receiver{
			Name: recv.Name(),
			Type: types.TypeString(recv.Type(), qualifier),
		}
	}

	// Params
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

	// Results
	results := sig.Results()
	for i := 0; i < results.Len(); i++ {
		r := results.At(i)
		info.Results = append(info.Results, &models.Field{
			Name:  r.Name(),
			Type:  types.TypeString(r.Type(), qualifier),
			Index: i,
		})
	}

	*funcs = append(*funcs, info)
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
