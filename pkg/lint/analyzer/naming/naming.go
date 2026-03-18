package naming

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
	"path/filepath"
	"strings"
)

var blacklist = map[string]string{
	"url":   "locator",
	"mcp":   "model_context",
	"dir":   "directory",
	"dirs":  "directories",
	"tx":    "t",
	"ctx":   "x",
	"param": "parameter",
	"msg":   "message",
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
	i.Preorder([]ast.Node{(*ast.Ident)(nil)}, func(n ast.Node) {
		ident := n.(*ast.Ident)
		file := filepath.Base(pass.Fset.File(ident.Pos()).Name())

		if skip[file] {
			return
		}

		if _, ok := pass.TypesInfo.Defs[ident]; !ok {
			return
		}

		check(pass, ident)
	})

	return nil, nil
}

func check(pass *analysis.Pass, ident *ast.Ident) {
	for _, segment := range segments(ident.Name) {
		if suggestion, ok := blacklist[segment]; ok {
			if !containsSegment(ident.Name, suggestion) {
				pass.Reportf(
					ident.Pos(),
					"use %q instead of %q in %s",
					suggestion,
					segment,
					ident.Name,
				)

				return
			}
		}
	}
}

func containsSegment(name string, target string) bool {
	for _, s := range segments(name) {
		if len(s) >= len(target) && s[:len(target)] == target {
			return true
		}
	}

	return false
}

func init() {
	var entries []string

	for k, v := range blacklist {
		entries = append(entries, fmt.Sprintf("%s -> %s", k, v))
	}

	Analyzer.Doc = fmt.Sprintf(
		"flags blacklisted identifier name segments: %v",
		entries,
	)
}
