package gofix

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

func findImportedPackage(
	p *packages.Package,
	path string,
) *types.Package {
	for _, i := range p.Types.Imports() {
		if i.Path() == path {
			return i
		}
	}

	return nil
}
