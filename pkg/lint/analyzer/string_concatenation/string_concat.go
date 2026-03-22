package string_concatenation

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:     "string_concat",
	Doc:      "flags string + concatenation; use fmt.Sprintf instead",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}
