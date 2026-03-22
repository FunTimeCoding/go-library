package struct_literal

import (
	"golang.org/x/tools/go/analysis"
	"os"
	"path/filepath"
	"strings"
)

func modulePrefix(p *analysis.Pass) string {
	if len(p.Files) == 0 {
		return ""
	}

	d := filepath.Dir(p.Fset.File(p.Files[0].Pos()).Name())

	for {
		content, e := os.ReadFile(filepath.Join(d, "go.mod"))

		if e == nil {
			for _, line := range strings.Split(string(content), "\n") {
				trimmed := strings.TrimSpace(line)

				if strings.HasPrefix(trimmed, "module ") {
					return strings.TrimSpace(
						strings.TrimPrefix(
							trimmed,
							"module ",
						),
					)
				}
			}
		}

		parent := filepath.Dir(d)

		if parent == d {
			return ""
		}

		d = parent
	}
}
