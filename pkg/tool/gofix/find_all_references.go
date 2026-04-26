package gofix

import (
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func findAllReferences(
	all []*packages.Package,
	o types.Object,
) []reference {
	var result []reference
	seen := make(map[token.Pos]bool)
	target := o.Pos()

	for _, p := range all {
		for ident, definition := range p.TypesInfo.Defs {
			if definition != nil && definition.Pos() == target && !seen[ident.Pos()] {
				seen[ident.Pos()] = true
				result = append(result, reference{ident: ident, p: p})
			}
		}

		for ident, use := range p.TypesInfo.Uses {
			if use.Pos() == target && !seen[ident.Pos()] {
				seen[ident.Pos()] = true
				result = append(result, reference{ident: ident, p: p})
			}
		}
	}

	return result
}
