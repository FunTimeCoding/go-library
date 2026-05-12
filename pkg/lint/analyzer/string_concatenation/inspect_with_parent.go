package string_concatenation

import (
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func inspectWithParent(
	file *ast.File,
	generated map[string]bool,
	p *packages.Package,
	results *output.Results,
) {
	var walk func(
		ast.Node,
		ast.Node,
	)
	walk = func(
		n ast.Node,
		parent ast.Node,
	) {
		if n == nil {
			return
		}

		if generated[p.Fset.Position(n.Pos()).Filename] {
			return
		}

		switch node := n.(type) {
		case *ast.BinaryExpr:
			if node.Op == token.ADD && !isParentConcatenation(parent) {
				checkBinary(p, results, node)
			}
		case *ast.AssignStmt:
			if node.Tok == token.ADD_ASSIGN {
				checkAssign(p, results, node)
			}
		}

		ast.Inspect(
			n,
			func(child ast.Node) bool {
				if child == n {
					return true
				}

				if child != nil {
					walk(child, n)
				}

				return child == nil
			},
		)
	}
	walk(file, nil)
}
