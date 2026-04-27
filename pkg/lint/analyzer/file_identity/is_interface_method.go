package file_identity

import (
	"go/types"
	"golang.org/x/tools/go/analysis"
)

func isInterfaceMethod(
	p *analysis.Pass,
	f *types.Func,
) bool {
	s, okay := f.Type().(*types.Signature)

	if !okay || s.Recv() == nil {
		return false
	}

	r := s.Recv().Type()

	if i, isPointer := r.(*types.Pointer); isPointer {
		r = i.Elem()
	}

	named, okay := r.(*types.Named)

	if !okay {
		return false
	}

	methodName := f.Name()

	for _, m := range p.Pkg.Imports() {
		scope := m.Scope()

		for _, name := range scope.Names() {
			typeName, okay := scope.Lookup(name).(*types.TypeName)

			if !okay {
				continue
			}

			face, okay := typeName.Type().Underlying().(*types.Interface)

			if !okay {
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
