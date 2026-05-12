package type_receiver

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
	receiverTypes := map[string]bool{}

	for _, file := range p.Syntax {
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
		return
	}

	reported := false

	for _, file := range p.Syntax {
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

				results.AddBlocked(
					p.Fset.Position(s.Pos()).Filename,
					"multiple types with receivers in one package; extract to subpackage",
				)
				reported = true

				return false
			},
		)
	}
}
