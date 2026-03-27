package naming

import (
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func findReferences(
	p *analysis.Pass,
	o types.Object,
) []*ast.Ident {
	var result []*ast.Ident

	for ident, definition := range p.TypesInfo.Defs {
		if definition == o {
			result = append(result, ident)
		}
	}

	for ident, use := range p.TypesInfo.Uses {
		if use == o {
			result = append(result, ident)
		}
	}

	return result
}
