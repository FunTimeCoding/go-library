package naming

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"path/filepath"
)

func run(p *analysis.Pass) (any, error) {
	skip := make(map[string]bool)

	for _, f := range p.Files {
		name := filepath.Base(p.Fset.File(f.Pos()).Name())

		if name == constant.GeneratedFile {
			skip[name] = true
		}
	}

	i := p.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	i.Preorder(
		[]ast.Node{(*ast.Ident)(nil)}, func(n ast.Node) {
			ident := n.(*ast.Ident)
			file := filepath.Base(p.Fset.File(ident.Pos()).Name())

			if skip[file] {
				return
			}

			o, isDefined := p.TypesInfo.Defs[ident]

			if !isDefined {
				return
			}

			check(p, ident, o)
		},
	)

	return nil, nil
}
