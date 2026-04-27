package naming

import (
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
)

var Analyzer = &analysis.Analyzer{
	Name:     "naming",
	Doc:      "flags blacklisted identifier name segments",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}
