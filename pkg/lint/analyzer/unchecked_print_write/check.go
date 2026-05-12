package unchecked_print_write

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	for _, file := range p.Syntax {
		if filepath.Base(p.Fset.File(file.Pos()).Name()) == constant.GeneratedFile {
			continue
		}

		if ast.IsGenerated(file) {
			continue
		}

		ast.Inspect(
			file,
			func(n ast.Node) bool {
				statement, okay := n.(*ast.ExprStmt)

				if !okay {
					return true
				}

				call, okay := statement.X.(*ast.CallExpr)

				if !okay {
					return true
				}

				checkCall(p, results, call)

				return true
			},
		)
	}
}
