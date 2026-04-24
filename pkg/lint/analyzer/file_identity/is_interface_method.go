package file_identity

import (
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func isInterfaceMethod(
	p *analysis.Pass,
	f *types.Func,
) bool {
	s, ok := f.Type().(*types.Signature)

	if !ok || s.Recv() == nil {
		return false
	}

	r := s.Recv().Type()

	if i, isPointer := r.(*types.Pointer); isPointer {
		r = i.Elem()
	}

	named, ok := r.(*types.Named)

	if !ok {
		return false
	}

	methodName := f.Name()

	for _, m := range p.Pkg.Imports() {
		scope := m.Scope()

		for _, name := range scope.Names() {
			typeName, ok := scope.Lookup(name).(*types.TypeName)

			if !ok {
				continue
			}

			face, ok := typeName.Type().Underlying().(*types.Interface)

			if !ok {
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
