package naming

import "go/types"

func isInterfaceMethod(
	o types.Object,
) bool {
	f, isFunction := o.(*types.Func)

	if !isFunction {
		return false
	}

	s, isSignature := f.Type().(*types.Signature)

	if !isSignature || s.Recv() == nil {
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

	for _, m := range f.Pkg().Imports() {
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
