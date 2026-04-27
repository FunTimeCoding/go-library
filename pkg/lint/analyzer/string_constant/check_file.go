package string_constant

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
)

func checkFile(
	p *analysis.Pass,
	file *ast.File,
	constants map[string]knownConstant,
) {
	ast.Inspect(
		file,
		func(n ast.Node) bool {
			call, okay := n.(*ast.CallExpr)

			if !okay {
				return true
			}

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

			return true
		},
	)
}
