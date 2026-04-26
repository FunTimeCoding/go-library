package string_constant

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"path/filepath"
)

func run(p *analysis.Pass) (any, error) {
	if len(p.Files) == 0 {
		return nil, nil
	}

	directory := filepath.Dir(p.Fset.File(p.Files[0].Pos()).Name())

	if filepath.Base(directory) == "constant" {
		return nil, nil
	}

	constants := collectConstants(directory)

	if len(constants) == 0 {
		return nil, nil
	}

	for _, file := range p.Files {
		name := filepath.Base(p.Fset.File(file.Pos()).Name())

		if name == constant.GeneratedFile {
			continue
		}

		if name == "constant.go" {
			continue
		}

		if ast.IsGenerated(file) {
			continue
		}

		checkFile(p, file, constants)
	}

	return nil, nil
}
