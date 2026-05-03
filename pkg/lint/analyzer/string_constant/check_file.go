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
			switch v := n.(type) {
			case *ast.CallExpr:
				checkCallArguments(p, v, constants)
			case *ast.IndexExpr:
				checkArgument(p, v.Index, constants)
			}

			return true
		},
	)
}
