package string_constant

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func checkCallArguments(
	p *analysis.Pass,
	call *ast.CallExpr,
	constants map[string]knownConstant,
) {
	skipIndex := -1

	if isAssertCall(p, call) {
		skipIndex = 1
	}

	for i, a := range call.Args {
		if i == skipIndex {
			continue
		}

		checkArgument(p, a, constants)
	}
}
