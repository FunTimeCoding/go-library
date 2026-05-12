package string_concatenation

import (
	"github.com/funtimecoding/go-library/pkg/constant"
	"github.com/funtimecoding/go-library/pkg/lint/output"
	"golang.org/x/tools/go/packages"
	"path/filepath"
)

func Check(
	p *packages.Package,
	results *output.Results,
) {
	generated := make(map[string]bool)

	for _, file := range p.Syntax {
		name := p.Fset.File(file.Pos()).Name()

		if filepath.Base(name) == constant.GeneratedFile {
			generated[name] = true
		}
	}

	for _, file := range p.Syntax {
		inspectWithParent(
			file,
			generated,
			p,
			results,
		)
	}
}
