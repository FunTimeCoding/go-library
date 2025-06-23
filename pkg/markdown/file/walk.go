package file

import "github.com/yuin/goldmark/ast"

func Walk(
	s *[]byte,
	n ast.Node,
) {
	if n == nil {
		return
	}

	PrintNode(s, n)

	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		Walk(s, c)
	}
}
