package runbook

import (
	"github.com/yuin/goldmark/ast"
	"strings"
)

func extractCode(
	source *[]byte,
	c *ast.FencedCodeBlock,
) string {
	var b strings.Builder

	for i := 0; i < c.Lines().Len(); i++ {
		l := c.Lines().At(i)
		b.Write(l.Value(*source))
	}

	return strings.TrimSpace(b.String())
}
