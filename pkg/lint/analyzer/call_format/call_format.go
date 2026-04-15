package call_format

import "golang.org/x/tools/go/analysis"

var Analyzer = &analysis.Analyzer{
	Name: "call_format",
	Doc:  "flags function calls where multiple arguments share a line",
	Run:  run,
}
