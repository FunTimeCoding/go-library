package forbidden_import

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"strings"
)

func run(p *analysis.Pass) (any, error) {
	for _, file := range p.Files {
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

				p.Reportf(i.Pos(), "%s", b.message)
			}
		}
	}

	return nil, nil
}
