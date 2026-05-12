package struct_literal

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	if p.Module == nil {
		return
	}

	module := p.Module.Path
	generated := make(map[string]bool)

	for _, file := range p.Syntax {
		name := p.Fset.File(file.Pos()).Name()

		if filepath.Base(name) == constant.GeneratedFile {
			generated[name] = true
		}
	}

	for _, file := range p.Syntax {
		ast.Inspect(
			file,
			func(n ast.Node) bool {
				if n == nil {
					return true
				}

				if generated[p.Fset.Position(n.Pos()).Filename] {
					return true
				}

				switch node := n.(type) {
				case *ast.UnaryExpr:
					if node.Op == token.AND {
						checkUnary(p, results, node, module)
					}
				case *ast.CallExpr:
					checkNew(p, results, node, module)
				}

				return true
			},
		)
	}
}
