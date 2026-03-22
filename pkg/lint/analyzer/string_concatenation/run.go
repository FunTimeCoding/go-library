package string_concatenation

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"path/filepath"
)

func run(p *analysis.Pass) (any, error) {
	generated := make(map[string]bool)

	for _, file := range p.Files {
		name := p.Fset.File(file.Pos()).Name()

		if filepath.Base(name) == constant.GeneratedFile {
			generated[name] = true
		}
	}

	i := p.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	i.WithStack(
		[]ast.Node{(*ast.BinaryExpr)(nil), (*ast.AssignStmt)(nil)},
		func(n ast.Node, push bool, stack []ast.Node) bool {
			if !push {
				return true
			}

			if generated[p.Fset.Position(n.Pos()).Filename] {
				return true
			}

			switch node := n.(type) {
			case *ast.BinaryExpr:
				if node.Op == token.ADD && !isNestedConcatenation(stack) {
					checkBinary(p, node)
				}
			case *ast.AssignStmt:
				if node.Tok == token.ADD_ASSIGN {
					checkAssign(p, node)
				}
			}

			return true
		},
	)

	return nil, nil
}
