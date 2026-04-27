package type_receiver

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"path/filepath"
)

func run(p *analysis.Pass) (any, error) {
	receiverTypes := map[string]bool{}

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
				f, okay := n.(*ast.FuncDecl)

				if !okay || f.Recv == nil {
					return true
				}

				name := receiverTypeName(f.Recv.List[0].Type)

				if name != "" {
					receiverTypes[name] = true
				}

				return true
			},
		)
	}

	if len(receiverTypes) < 2 {
		return nil, nil
	}

	reported := false

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
				if reported {
					return false
				}

				s, okay := n.(*ast.TypeSpec)

				if !okay {
					return true
				}

				if !receiverTypes[s.Name.Name] {
					return true
				}

				p.Reportf(
					s.Pos(),
					"multiple types with receivers in one package; extract to subpackage",
				)
				reported = true

				return false
			},
		)
	}

	return nil, nil
}
