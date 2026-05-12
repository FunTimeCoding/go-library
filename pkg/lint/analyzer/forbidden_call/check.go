package forbidden_call

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
	generated := make(map[string]bool)

	for _, file := range p.Syntax {
		name := p.Fset.File(file.Pos()).Name()

		if filepath.Base(name) == constant.GeneratedFile || ast.IsGenerated(file) {
			generated[name] = true
		}
	}

	for _, file := range p.Syntax {
		ast.Inspect(
			file,
			func(n ast.Node) bool {
				call, okay := n.(*ast.CallExpr)

				if !okay {
					return true
				}

				if generated[p.Fset.Position(call.Pos()).Filename] {
					return true
				}

				checkCall(p, results, call)

				return true
			},
		)
	}
}
