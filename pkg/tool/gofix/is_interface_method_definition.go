package gofix

import "go/types"

func isInterfaceMethodDefinition(o types.Object) bool {
	f, isFunction := o.(*types.Func)

	if !isFunction {
		return false
	}

	s, isSignature := f.Type().(*types.Signature)

	if !isSignature || s.Recv() == nil {
		return false
	}

	_, isInterface := s.Recv().Type().Underlying().(*types.Interface)

	return isInterface
}
