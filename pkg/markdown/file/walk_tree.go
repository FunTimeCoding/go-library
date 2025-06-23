package file

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/markdown/file/flat"
	"github.com/yuin/goldmark/ast"
)

func WalkTree(
	s *[]byte,
	n ast.Node,
	l *flat.Flat,
) *flat.Flat {
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

				l.Add(NodeValue(s, n))
				PrintNode(s, n)

				if false {
					PrintKind(s, n)
				}

				return ast.WalkContinue, nil
			},
		),
	)

	return l
}
