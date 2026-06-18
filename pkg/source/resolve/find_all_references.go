package resolve

import (
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func FindAllReferences(
	all []*packages.Package,
	o types.Object,
) []Reference {
	var result []Reference
	seen := make(map[token.Pos]bool)
	target := o.Pos()

	for _, p := range all {
		for ident, definition := range p.TypesInfo.Defs {
			if definition != nil && definition.Pos() == target && !seen[ident.Pos()] {
				seen[ident.Pos()] = true
				result = append(result, Reference{Ident: ident, Package: p})
			}
		}

		for ident, use := range p.TypesInfo.Uses {
			if use.Pos() == target && !seen[ident.Pos()] {
				seen[ident.Pos()] = true
				result = append(result, Reference{Ident: ident, Package: p})
			}
		}
	}

	return result
}
