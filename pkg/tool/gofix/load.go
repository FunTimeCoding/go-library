package gofix

import (
	"fmt"
	"go/token"
	"golang.org/x/tools/go/packages"
	"os"
)

func load(
	fileSet *token.FileSet,
	patterns []string,
) []*packages.Package {
	configuration := &packages.Config{
		Mode:  packages.LoadSyntax,
		Fset:  fileSet,
		Tests: true,
	}
	result, e := packages.Load(configuration, patterns...)

	if e != nil {
		fmt.Fprintf(os.Stderr, "load: %s\n", e)
		os.Exit(1)
	}

	var hasError bool

	for _, p := range result {
		for _, e := range p.Errors {
			fmt.Fprintf(os.Stderr, "%s: %s\n", p.PkgPath, e)
			hasError = true
		}
	}

	if hasError {
		os.Exit(1)
	}

	return result
}
