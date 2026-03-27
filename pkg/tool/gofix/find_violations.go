package gofix

import (
	"go/token"
	"go/types"
	"golang.org/x/tools/go/packages"
	"path/filepath"
	"strings"
)

func findViolations(all []*packages.Package) []violation {
	var result []violation
	seen := make(map[token.Pos]bool)

	for _, p := range all {
		for ident, o := range p.TypesInfo.Defs {
			if o == nil {
				continue
			}

			if seen[o.Pos()] {
				continue
			}

			if !ident.IsExported() {
				continue
			}

			file := filepath.Base(p.Fset.File(ident.Pos()).Name())

			if strings.HasSuffix(file, "_gen.go") {
				continue
			}

			if _, isVariable := o.(*types.Var); isVariable {
				continue
			}

			v := checkNaming(ident, o)

			if v != nil {
				seen[o.Pos()] = true
				result = append(result, *v)
			}
		}
	}

	return result
}
