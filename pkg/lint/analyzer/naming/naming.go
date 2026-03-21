package naming

import (
	"fmt"
	"go/ast"
	"go/types"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"path/filepath"
	"strings"
)

// suggestions maps a banned name segment to acceptable alternatives.
// Single-character alternatives apply only in variable contexts
// (local variables, parameters, struct fields).
var suggestions = map[string][]string{
	"url":   {"l", "locator"},
	"mcp":   {"c", "model_context"},
	"dir":   {"d", "directory"},
	"dirs":  {"directories"},
	"tx":    {"t"}, // transaction is too long; t suffices
	"ctx":   {"x"}, // context collides with the standard library package name
	"param": {"parameter"},
	"msg":   {"m", "message"},
	"req":   {"r", "request"},
	"doc":   {"d", "document"},
	"config": {
		"c",
		"configuration",
	}, // configuration is long but beats a vague abbreviation in non-variable contexts; often replaceable with something more specific like setting
	"cfg": {"c", "configuration"},
	"llm": {"m", "model"},
	"tmp": {"t"}, // temporary is too long; t suffices
}

// noSuggestion contains banned noise words with no prescribed replacement.
// They are flagged regardless of identifier kind.
var noSuggestion = map[string]bool{
	"handler": true, // too vague; use a name that describes what it handles
	"data":    true, // too vague; everything is data
	"info":    true, // too vague; everything is info
}

var Analyzer = &analysis.Analyzer{
	Name:     "naming",
	Doc:      "flags blacklisted identifier name segments",
	Requires: []*analysis.Analyzer{inspect.Analyzer},
	Run:      run,
}

func run(pass *analysis.Pass) (any, error) {
	skip := make(map[string]bool)

	for _, f := range pass.Files {
		name := filepath.Base(pass.Fset.File(f.Pos()).Name())

		if name == "generated.go" || strings.HasPrefix(name, "generated_") {
			skip[name] = true
		}
	}

	i := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	i.Preorder(
		[]ast.Node{(*ast.Ident)(nil)}, func(n ast.Node) {
			ident := n.(*ast.Ident)
			file := filepath.Base(pass.Fset.File(ident.Pos()).Name())

			if skip[file] {
				return
			}

			obj, ok := pass.TypesInfo.Defs[ident]

			if !ok {
				return
			}

			check(pass, ident, obj)
		},
	)

	return nil, nil
}

func check(
	pass *analysis.Pass,
	ident *ast.Ident,
	obj types.Object,
) {
	_, isVar := obj.(*types.Var)

	for _, segment := range segments(ident.Name) {
		if noSuggestion[segment] {
			pass.Reportf(
				ident.Pos(),
				"avoid %q in name, use a more specific term",
				segment,
			)

			return
		}

		alts, ok := suggestions[segment]

		if !ok {
			continue
		}

		var applicable []string

		for _, alt := range alts {
			if len(alt) > 1 || isVar {
				applicable = append(applicable, alt)
			}
		}

		if len(applicable) == 0 {
			pass.Reportf(
				ident.Pos(),
				"avoid %q in name, use a more specific term",
				segment,
			)

			return
		}

		for _, alt := range applicable {
			if containsSegment(ident.Name, alt) {
				return
			}
		}

		pass.Reportf(
			ident.Pos(),
			"%s",
			formatMessage(applicable, segment, ident.Name),
		)

		return
	}
}

func formatMessage(
	applicable []string,
	segment, name string,
) string {
	if len(applicable) == 1 {
		return fmt.Sprintf(
			"use %q instead of %q in %s",
			applicable[0],
			segment,
			name,
		)
	}

	quoted := make([]string, len(applicable))

	for i, a := range applicable {
		quoted[i] = fmt.Sprintf("%q", a)
	}

	return fmt.Sprintf(
		"use %s instead of %q in %s",
		strings.Join(quoted, " or "),
		segment,
		name,
	)
}

func containsSegment(
	name string,
	target string,
) bool {
	for _, s := range segments(name) {
		if len(target) == 1 {
			if s == target {
				return true
			}
		} else {
			if len(s) >= len(target) && s[:len(target)] == target {
				return true
			}
		}
	}

	return false
}

func init() {
	var entries []string

	for k, v := range suggestions {
		entries = append(
			entries,
			fmt.Sprintf("%s -> %s", k, strings.Join(v, " or ")),
		)
	}

	for k := range noSuggestion {
		entries = append(entries, fmt.Sprintf("%s (avoid)", k))
	}

	Analyzer.Doc = fmt.Sprintf(
		"flags blacklisted identifier name segments: %v",
		entries,
	)
}
