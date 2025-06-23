package file

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/yuin/goldmark/ast"
)

func WalkTree(
	s *[]byte,
	n ast.Node,
) {
	errors.PanicOnError(
		ast.Walk(
			n,
			func(
				n ast.Node,
				entering bool,
			) (ast.WalkStatus, error) {
				if !entering {
					return ast.WalkContinue, nil
				}

				PrintNode(s, n)

				if false {
					if n.Kind() == ast.KindText {
						textNode := n.(*ast.Text)
						segment := textNode.Segment
						fmt.Printf(
							"Text kind: %s\n",
							string(segment.Value(*s)),
						)
					}
				}

				return ast.WalkContinue, nil
			},
		),
	)
}
