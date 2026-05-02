package gofix

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"go/token"
	"golang.org/x/tools/go/packages"
	"os"
)

func load(
	fileSet *token.FileSet,
	directory string,
	patterns []string,
) []*packages.Package {
	configuration := &packages.Config{
		Mode:  packages.LoadSyntax,
		Fset:  fileSet,
		Dir:   directory,
		Tests: true,
	}
	result, e := packages.Load(configuration, patterns...)

	if e != nil {
		errors.Printf("load: %s\n", e)
		os.Exit(1)
	}

	var hasError bool

	for _, p := range result {
		for _, e := range p.Errors {
			errors.Printf("%s: %s\n", p.PkgPath, e)
			hasError = true
		}
	}

	if hasError {
		os.Exit(1)
	}

	return result
}
