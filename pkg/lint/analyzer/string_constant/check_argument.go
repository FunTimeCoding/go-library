package string_constant

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/ast"
	"go/token"
	"golang.org/x/tools/go/packages"
	"strings"
)

func checkArgument(
	p *packages.Package,
	results *output.Results,
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

	results.AddBlocked(
		p.Fset.Position(l.Pos()).Filename,
		fmt.Sprintf(
			"string literal %q has constant %s.%s",
			value,
			c.packageName,
			c.name,
		),
	)
}
