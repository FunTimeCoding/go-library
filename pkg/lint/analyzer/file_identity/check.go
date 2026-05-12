package file_identity

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	for _, file := range p.Syntax {
		name := filepath.Base(p.Fset.File(file.Pos()).Name())

		if skip(name) {
			continue
		}

		if name == constant.GeneratedFile {
			continue
		}

		if ast.IsGenerated(file) {
			continue
		}

		checkFile(p, results, file, name)
	}
}
