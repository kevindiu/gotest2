package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kevindiu/gotest2/internal/app"
)

func main() {
	var (
		all          = flag.Bool("all", false, "generate tests for all functions and methods")
		exported     = flag.Bool("exported", true, "generate tests for exported functions and methods. Usage: -exported=false to include unexported")
		templatePath = flag.String("template", "", "path to custom template file")
		parallel     = flag.Bool("parallel", true, "generate parallel tests")
		fuzz         = flag.Bool("fuzz", false, "generate fuzz tests (where applicable)")
		benchmark    = flag.Bool("benchmark", false, "generate benchmark functions")
		tests        = flag.Bool("tests", true, "generate standard unit tests")
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

	cfg := app.Config{
		All:          *all,
		Exported:     *exported,
		TemplatePath: *templatePath,
		Parallel:     *parallel,
		Fuzz:         *fuzz,
		Benchmark:    *benchmark,
		Tests:        *tests,
	}

	if err := app.Run(files, cfg); err != nil {
		log.Fatal(err)
	}
}
