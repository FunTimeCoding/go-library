package file

import "github.com/yuin/goldmark/ast"

func Value(
	s *[]byte,
	n ast.Node,
) string {
	return string(n.Lines().Value(*s))
}
