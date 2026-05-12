package anonymous_struct

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"strings"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	for _, file := range p.Syntax {
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

				checkStruct(p, results, s)

				return true
			},
		)
	}
}
