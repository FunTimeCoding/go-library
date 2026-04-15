package defer_close

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
				d, ok := n.(*ast.DeferStmt)

				if !ok {
					return true
				}

				check(p, d)

				return true
			},
		)
	}

	return nil, nil
}
