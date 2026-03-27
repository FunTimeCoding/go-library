package gofix

import (
	"go/types"
	"golang.org/x/tools/go/packages"
)

func isInterfaceMethod(
	p *packages.Package,
	o types.Object,
) bool {
	f, isFunc := o.(*types.Func)

	if !isFunc {
		return false
	}

	s, isSig := f.Type().(*types.Signature)

	if !isSig || s.Recv() == nil {
		return false
	}

	r := s.Recv().Type()

	if i, isPointer := r.(*types.Pointer); isPointer {
		r = i.Elem()
	}

	named, isNamed := r.(*types.Named)

	if !isNamed {
		return false
	}

	methodName := f.Name()

	for _, m := range p.Types.Imports() {
		scope := m.Scope()

		for _, name := range scope.Names() {
			typeName, isTypeName := scope.Lookup(name).(*types.TypeName)

			if !isTypeName {
				continue
			}

			face, isFace := typeName.Type().Underlying().(*types.Interface)

			if !isFace {
				continue
			}

			for i := range face.NumMethods() {
				if face.Method(i).Name() != methodName {
					continue
				}

				if types.Implements(named, face) ||
					types.Implements(types.NewPointer(named), face) {
					return true
				}
			}
		}
	}

	return false
}
