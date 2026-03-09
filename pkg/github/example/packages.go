package example

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/constant"
)

func Packages() {
	c := github.NewEnvironment()
	f := constant.Format

	for _, p := range c.Packages(constant.LibraryNamespace) {
		fmt.Printf("Package: %s\n", p.Format(f))

		for _, v := range c.PackageVersions(p.Name) {
			fmt.Printf("  Image: %s\n", v.Format(f))
		}
	}
}
