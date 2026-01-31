package runbook

import (
	"github.com/yuin/goldmark/ast"
	"strings"
)

func extractText(
	source *[]byte,
	n ast.Node,
) string {
	var b strings.Builder

	for c := n.FirstChild(); c != nil; c = c.NextSibling() {
		if c.Kind() == ast.KindText {
			b.Write(c.(*ast.Text).Segment.Value(*source))
		}
	}

	return b.String()
}
