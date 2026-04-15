package call_format

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:     "call_format",
	Doc:      "flags function calls where multiple arguments share a line",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}
