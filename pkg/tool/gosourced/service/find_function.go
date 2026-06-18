package service

import (
	"fmt"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func findFunction(
	p *packages.Package,
	symbol string,
) (types.Object, *packages.Package, error) {
	o := p.Types.Scope().Lookup(symbol)

	if o == nil {
		return nil, nil, fmt.Errorf(
			"symbol %s not found in %s",
			symbol,
			p.PkgPath,
		)
	}

	return o, p, nil
}
