package goanalyze

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/lint/analyzer/variable_naming"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func findVariableNamingEdits(all []*packages.Package) []edit {
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
					fn, ok := n.(*ast.FuncDecl)

					if !ok || fn.Body == nil {
						return true
					}

					renames := variable_naming.ComputeRenames(p, fn)

					for _, rename := range renames {
						if seen[rename.Object.Pos()] {
							continue
						}

						seen[rename.Object.Pos()] = true
						fmt.Printf(
							"%s → %s\n",
							rename.Object.Name(),
							rename.NewName,
						)
						references := findAllReferences(
							all,
							rename.Object,
						)

						for _, r := range references {
							result = append(
								result,
								edit{
									position: r.ident.Pos(),
									end:      r.ident.End(),
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
