package file

import (
	"fmt"
	"github.com/yuin/goldmark/ast"
)

func PrintKind(
	s *[]byte,
	n ast.Node,
) {
	if n.Kind() == ast.KindText {
		t := n.(*ast.Text)
		e := t.Segment
		fmt.Printf("Text kind: %s\n", string(e.Value(*s)))
	}
}
