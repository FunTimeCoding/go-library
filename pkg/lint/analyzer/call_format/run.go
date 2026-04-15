package call_format

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"go/ast"
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
	i.Preorder(
		[]ast.Node{(*ast.CallExpr)(nil)},
		func(n ast.Node) {
			if generated[p.Fset.Position(n.Pos()).Filename] {
				return
			}

			check(p, n.(*ast.CallExpr))
		},
	)

	return nil, nil
}
