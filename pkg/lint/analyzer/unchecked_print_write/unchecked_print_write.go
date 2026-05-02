package unchecked_print_write

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "unchecked_print_write",
	Doc:  "flags fmt.Fprintf calls where the error return is discarded; use writer.Print, errors.Printf, or check the error",
	Run:  run,
}
