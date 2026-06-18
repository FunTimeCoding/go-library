package service

import (
	"fmt"
	"go/types"
	"golang.org/x/tools/go/packages"
)

func findMethod(
	p *packages.Package,
	symbol string,
	receiver string,
) (types.Object, *packages.Package, error) {
	o := p.Types.Scope().Lookup(receiver)

	if o == nil {
		return nil, nil, fmt.Errorf(
			"receiver type %s not found in %s",
			receiver,
			p.PkgPath,
		)
	}

	named, okay := o.Type().(*types.Named)

	if !okay {
		return nil, nil, fmt.Errorf(
			"%s is not a named type in %s",
			receiver,
			p.PkgPath,
		)
	}

	for i := range named.NumMethods() {
		m := named.Method(i)

		if m.Name() == symbol {
			return m, p, nil
		}
	}

	return nil, nil, fmt.Errorf(
		"method %s not found on %s in %s",
		symbol,
		receiver,
		p.PkgPath,
	)
}
