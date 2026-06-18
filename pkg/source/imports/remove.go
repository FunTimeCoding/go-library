package imports

import (
	"go/ast"
	"go/token"
)

func Remove(
	file *ast.File,
	spec *ast.ImportSpec,
) {
	importPath := spec.Path.Value

	for _, group := range file.Decls {
		declaration, okay := group.(*ast.GenDecl)

		if !okay || declaration.Tok != token.IMPORT {
			continue
		}

		for j, s := range declaration.Specs {
			candidate := s.(*ast.ImportSpec)

			if candidate.Path.Value != importPath {
				continue
			}

			if sameAlias(candidate, spec) {
				declaration.Specs = append(
					declaration.Specs[:j],
					declaration.Specs[j+1:]...,
				)

				return
			}
		}
	}
}
