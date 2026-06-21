package goanalyze

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/source/resolve"
	"golang.org/x/tools/go/packages"
)

func load(patterns []string) []*packages.Package {
	result, _, e := resolve.LoadPackages("", patterns...)
	errors.PanicOnError(e)

	return result
}
