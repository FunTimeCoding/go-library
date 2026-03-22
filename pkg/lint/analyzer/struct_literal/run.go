package struct_literal

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
	module := modulePrefix(p)

	if module == "" {
		return nil, nil
	}

	generated := make(map[string]bool)

	for _, file := range p.Files {
		name := p.Fset.File(file.Pos()).Name()

		if filepath.Base(name) == constant.GeneratedFile {
			generated[name] = true
		}
	}

	i := p.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	i.Preorder(
		[]ast.Node{(*ast.UnaryExpr)(nil), (*ast.CallExpr)(nil)},
		func(n ast.Node) {
			if generated[p.Fset.Position(n.Pos()).Filename] {
				return
			}

			switch node := n.(type) {
			case *ast.UnaryExpr:
				if node.Op == token.AND {
					checkUnary(p, node, module)
				}
			case *ast.CallExpr:
				checkNew(p, node, module)
			}
		},
	)

	return nil, nil
}
