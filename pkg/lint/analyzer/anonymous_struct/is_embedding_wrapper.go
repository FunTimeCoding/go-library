package anonymous_struct

import "go/ast"

func isEmbeddingWrapper(s *ast.StructType) bool {
	if len(s.Fields.List) != 1 {
		return false
	}

	field := s.Fields.List[0]

	if len(field.Names) != 0 {
		return false
	}

	_, isPointer := field.Type.(*ast.StarExpr)

	return isPointer
}
