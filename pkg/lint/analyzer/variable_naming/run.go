package variable_naming

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"path/filepath"
)

func run(p *analysis.Pass) (any, error) {
	for _, file := range p.Files {
		if filepath.Base(p.Fset.File(file.Pos()).Name()) == constant.GeneratedFile {
			continue
		}

		if ast.IsGenerated(file) {
			continue
		}

		ast.Inspect(
			file,
			func(n ast.Node) bool {
				fn, ok := n.(*ast.FuncDecl)

				if !ok || fn.Body == nil {
					return true
				}

				checkFunction(p, fn)

				return true
			},
		)
	}

	return nil, nil
}
