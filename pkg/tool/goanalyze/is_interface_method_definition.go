package goanalyze

import "go/types"

func isInterfaceMethodDefinition(o types.Object) bool {
	f, isFunc := o.(*types.Func)

	if !isFunc {
		return false
	}

	s, isSig := f.Type().(*types.Signature)

	if !isSig || s.Recv() == nil {
		return false
	}

	_, isInterface := s.Recv().Type().Underlying().(*types.Interface)

	return isInterface
}
