package string_constant

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
	if len(p.Syntax) == 0 {
		return
	}

	directory := filepath.Dir(p.Fset.File(p.Syntax[0].Pos()).Name())

	if filepath.Base(directory) == "constant" {
		return
	}

	constants := collectConstants(directory)

	if len(constants) == 0 {
		return
	}

	for _, file := range p.Syntax {
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

		checkFile(p, results, file, constants)
	}
}
