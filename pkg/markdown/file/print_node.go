package file

import (
	"fmt"
	"github.com/yuin/goldmark/ast"
)

func PrintNode(
	s *[]byte,
	n ast.Node,
) {
	switch o := n.(type) {
	case *ast.Document:
		v := Value(s, o)

		if v == "" {
			return
		}

		fmt.Printf("Document: %s\n", v)
	case *ast.Heading:
		fmt.Printf("Heading: %s\n", Value(s, o))
	case *ast.Paragraph:
		fmt.Printf("Paragraph: %s\n", Value(s, o))
	case *ast.Text:
		fmt.Printf("Text: %s\n", string(o.Value(*s)))
	case *ast.FencedCodeBlock:
		fmt.Printf("FencedCodeBlock: %s\n", Value(s, o))
	default:
		fmt.Printf("Unknown node type: %T\n", o)
	}
}
