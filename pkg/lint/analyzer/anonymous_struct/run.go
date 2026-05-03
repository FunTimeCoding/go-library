package anonymous_struct

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"path/filepath"
	"strings"
)

func run(p *analysis.Pass) (any, error) {
	for _, file := range p.Files {
		name := filepath.Base(p.Fset.File(file.Pos()).Name())

		if name == constant.GeneratedFile {
			continue
		}

		if strings.HasSuffix(name, "_test.go") {
			continue
		}

		if ast.IsGenerated(file) {
			continue
		}

		named := collectNamedStructPositions(file)
		ast.Inspect(
			file,
			func(n ast.Node) bool {
				s, okay := n.(*ast.StructType)

				if !okay {
					return true
				}

				if named[s.Pos()] {
					return true
				}

				check(p, s)

				return true
			},
		)
	}

	return nil, nil
}
