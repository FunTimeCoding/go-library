package string_constant

import (
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/analysis"
	"strings"
)

func checkArgument(
	p *analysis.Pass,
	e ast.Expr,
	constants map[string]knownConstant,
) {
	l, okay := e.(*ast.BasicLit)

	if !okay || l.Kind != token.STRING {
		return
	}

	value := strings.Trim(l.Value, "\"")

	if value == "" {
		return
	}

	c, exists := constants[value]

	if !exists {
		return
	}

	p.Reportf(
		l.Pos(),
		"string literal %q has constant %s.%s",
		value,
		c.packageName,
		c.name,
	)
}
