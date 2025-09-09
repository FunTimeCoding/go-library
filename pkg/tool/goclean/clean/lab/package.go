package lab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/packages"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
)

func Package(
	c *gitlab.Client,
	p *project.Project,
) {
	projectPackages := c.Packages(p.Identifier, false)

	if len(projectPackages) > 0 {
		latest := packages.Latest(projectPackages)

		for _, a := range projectPackages {
			var isLatest bool

			for _, inner := range latest {
				if inner.Name == a.Name &&
					inner.Version == a.Version {
					isLatest = true
				}
			}

			if isLatest {
				continue
			}

			fmt.Printf(
				"Delete package: %s %s\n",
				a.Name,
				a.Version,
			)
			c.DeletePackage(p.Identifier, a.ID)
		}
	}
}
