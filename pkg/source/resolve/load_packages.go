package resolve

import (
	"fmt"
	"go/token"
	"golang.org/x/tools/go/packages"
)

func LoadPackages(
	directory string,
	patterns ...string,
) ([]*packages.Package, *token.FileSet, error) {
	fileSet := token.NewFileSet()
	configuration := &packages.Config{
		Mode:  packages.LoadSyntax | packages.NeedModule,
		Fset:  fileSet,
		Dir:   directory,
		Tests: true,
	}
	result, e := packages.Load(configuration, patterns...)

	if e != nil {
		return nil, nil, fmt.Errorf("load packages: %w", e)
	}

	for _, p := range result {
		for _, e := range p.Errors {
			return nil, nil, fmt.Errorf("%s: %s", p.PkgPath, e)
		}
	}

	return result, fileSet, nil
}
