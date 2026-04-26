package string_constant

import (
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func parseConstants(
	result map[string]knownConstant,
	path string,
	p string,
	distance int,
) {
	set := token.NewFileSet()
	file, e := parser.ParseFile(set, path, nil, 0)

	if e != nil {
		return
	}

	if p == "" {
		p = file.Name.Name
	}

	for _, d := range file.Decls {
		g, ok := d.(*ast.GenDecl)

		if !ok || g.Tok != token.CONST {
			continue
		}

		for _, s := range g.Specs {
			v, ok := s.(*ast.ValueSpec)

			if !ok {
				continue
			}

			for i, name := range v.Names {
				if i >= len(v.Values) {
					continue
				}

				l, ok := v.Values[i].(*ast.BasicLit)

				if !ok || l.Kind != token.STRING {
					continue
				}

				value := strings.Trim(l.Value, "\"")
				existing, exists := result[value]

				if !exists || distance < existing.distance {
					result[value] = knownConstant{
						name:        name.Name,
						packageName: p,
						value:       value,
						distance:    distance,
					}
				}
			}
		}
	}
}
