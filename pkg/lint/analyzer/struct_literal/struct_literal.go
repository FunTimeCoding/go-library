package struct_literal

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:     "struct_literal",
	Doc:      "flags &pkg.Struct{} and new(pkg.Struct) for owned packages; use a constructor function instead",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}
