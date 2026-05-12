package parse

import (
	"go/ast"
	"go/token"
)

func FindVariable(
	f *ast.File,
	name string,
) *ast.ValueSpec {
	for _, d := range f.Decls {
		g, okay := d.(*ast.GenDecl)

		if !okay || g.Tok != token.VAR {
			continue
		}

		for _, s := range g.Specs {
			v, okay := s.(*ast.ValueSpec)

			if !okay {
				continue
			}

			for _, n := range v.Names {
				if n.Name == name {
					return v
				}
			}
		}
	}

	return nil
}
