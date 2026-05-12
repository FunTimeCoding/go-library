package forbidden_import

import (
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"strings"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	for _, file := range p.Syntax {
		if ast.IsGenerated(file) {
			continue
		}

		for _, i := range file.Imports {
			path := strings.Trim(i.Path.Value, `"`)

			for key, b := range banned {
				if b.substring {
					if !strings.Contains(path, key) {
						continue
					}
				} else {
					if path != key {
						continue
					}
				}

				results.AddBlocked(
					p.Fset.Position(i.Pos()).Filename,
					b.message,
				)
			}
		}
	}
}
