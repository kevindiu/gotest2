package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/kevindiu/gotest2/internal/generator"
	"github.com/kevindiu/gotest2/internal/models"
	"github.com/kevindiu/gotest2/internal/parser"
)

func main() {
	var (
		all          = flag.Bool("all", false, "generate tests for all functions and methods")
		exported     = flag.Bool("exported", true, "generate tests for exported functions and methods. Usage: -exported=false to include unexported")
		templatePath = flag.String("template", "", "path to custom template file")
		parallel     = flag.Bool("parallel", true, "generate parallel tests")
		fuzz         = flag.Bool("fuzz", false, "generate fuzz tests (where applicable)")
		// TODO: Support excluding specific functions
	)
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "\tgotest2 [flags] [files]\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	files := flag.Args()
	if len(files) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	// Parse all arguments as patterns
	fmt.Printf("Processing %d patterns: %v\n", len(files), files)
	results, err := parser.Parse(files)
	if err != nil {
		log.Fatalf("Failed to parse packages: %v", err)
	}

	for _, result := range results {
		fmt.Printf("Processing %s (%s)...\n", result.Path, result.PackageName)

		if len(result.Functions) == 0 {
			fmt.Printf("No functions found to generate tests for in %s\n", result.Path)
			continue
		}

		// Check for existing tests to avoid overwriting them
		absPath := result.Path
		ext := filepath.Ext(absPath)
		base := absPath[:len(absPath)-len(ext)]
		testPath := base + "_test.go"

		var existingTests map[string]string
		if _, err := os.Stat(testPath); err == nil {
			existingTests, err = parser.ParseTests(testPath)
			if err != nil {
				log.Printf("Warning: failed to parse existing tests in %s: %v", testPath, err)
				// Continue anyway
			}
		}

		// Filter based on flags
		var targetFuncs []*models.FunctionInfo
		for _, fn := range result.Functions {
			if *all {
				targetFuncs = append(targetFuncs, fn)
				continue
			}
			if *exported && !fn.IsExported {
				continue
			}
			targetFuncs = append(targetFuncs, fn)
		}

		if len(targetFuncs) == 0 {
			continue
		}

		fmt.Printf("Generating tests for %d functions in %s\n", len(targetFuncs), result.Path)

		// Note: We use result.PackageName for the generated file
		code, err := generator.Generate(targetFuncs, result.PackageName, *templatePath, *parallel, *fuzz, existingTests)
		if err != nil {
			log.Printf("Failed to generate tests for %s: %v", result.Path, err)
			continue
		}

		if err := generator.WriteFile(absPath, code); err != nil {
			log.Printf("Failed to write test file: %v", err)
			continue
		}
	}
}
