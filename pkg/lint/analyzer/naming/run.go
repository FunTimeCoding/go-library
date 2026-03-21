package naming

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"path/filepath"
	"strings"
)

func run(pass *analysis.Pass) (any, error) {
	skip := make(map[string]bool)

	for _, f := range pass.Files {
		name := filepath.Base(pass.Fset.File(f.Pos()).Name())

		if name == "generated.go" || strings.HasPrefix(name, "generated_") {
			skip[name] = true
		}
	}

	i := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	i.Preorder(
		[]ast.Node{(*ast.Ident)(nil)}, func(n ast.Node) {
			ident := n.(*ast.Ident)
			file := filepath.Base(pass.Fset.File(ident.Pos()).Name())

			if skip[file] {
				return
			}

			obj, ok := pass.TypesInfo.Defs[ident]

			if !ok {
				return
			}

			check(pass, ident, obj)
		},
	)

	return nil, nil
}
