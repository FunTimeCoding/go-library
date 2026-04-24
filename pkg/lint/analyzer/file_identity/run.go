package file_identity

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"path/filepath"
)

func run(p *analysis.Pass) (any, error) {
	for _, file := range p.Files {
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

		checkFile(p, file, name)
	}

	return nil, nil
}
