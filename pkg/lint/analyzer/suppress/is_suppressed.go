package suppress

import (
	"fmt"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"strings"
)

// IsSuppressed reports whether node is suppressed by a goanalyze:ignore comment
// on the same line. key is the analyzer name (e.g. "string_concat").
// A bare "goanalyze:ignore" suppresses all analyzers; "goanalyze:ignore <key>"
// suppresses only the named analyzer.
func IsSuppressed(
	p *analysis.Pass,
	n ast.Node,
	key string,
) bool {
	position := p.Fset.Position(n.Pos())

	for _, file := range p.Files {
		if p.Fset.File(file.Pos()).Name() != position.Filename {
			continue
		}

		for _, group := range file.Comments {
			for _, comment := range group.List {
				if p.Fset.Position(comment.Pos()).Line != position.Line {
					continue
				}

				text := strings.TrimSpace(
					strings.TrimPrefix(comment.Text, "//"),
				)

				if text == "goanalyze:ignore" || text == fmt.Sprintf(
					"goanalyze:ignore %s",
					key,
				) {
					return true
				}
			}
		}
	}

	return false
}
