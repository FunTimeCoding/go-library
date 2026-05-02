package gofix

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/variable_naming"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func findVariableNamingEdits(
	fileSet *token.FileSet,
	all []*packages.Package,
	r *results,
) []edit {
	var result []edit
	seen := make(map[token.Pos]bool)

	for _, p := range all {
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

					if !okay || f.Body == nil {
						return true
					}

					renames := variable_naming.ComputeRenames(p, f)

					for _, rename := range renames {
						if seen[rename.Object.Pos()] {
							continue
						}

						seen[rename.Object.Pos()] = true
						path := fileSet.File(rename.Object.Pos()).Name()
						r.add(
							path,
							fmt.Sprintf(
								"renamed %s → %s",
								rename.Object.Name(),
								rename.NewName,
							),
						)
						references := findAllReferences(
							all,
							rename.Object,
						)

						for _, ref := range references {
							result = append(
								result,
								edit{
									position: ref.ident.Pos(),
									end:      ref.ident.End(),
									newText:  rename.NewName,
								},
							)
						}
					}

					return true
				},
			)
		}
	}

	return result
}
