package hub

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/image"
)

func Image(
	c *github.Client,
	namespace string,
	repository string,
) {
	for _, p := range c.Packages(namespace) {
		if p.Repository != repository {
			continue
		}

		for _, v := range c.PackageVersions(p.Name) {
			if image.HasLatest(v) {
				continue
			}

			fmt.Printf("Delete image: %s@%s\n", p.Name, v.Digest)
			c.DeletePackageVersion(p.Name, v.Identifier)
		}
	}
}
