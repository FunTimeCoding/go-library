package lab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/constant"
	"github.com/funtimecoding/go-library/pkg/gitlab/image"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"strings"
)

func Registry(
	c *gitlab.Client,
	p *project.Project,
) {
	for _, r := range c.RegistryRepositories(
		p.Identifier,
		false,
	) {
		images := c.Images(p.Identifier, r.ID)

		if len(images) == 0 {
			continue
		}

		if strings.HasSuffix(r.Name, "/cache") {
			for _, i := range images {
				fmt.Printf("Image: %s\n", i.Name)
				c.DeleteImage(p.Identifier, r.ID, i.Name)
			}

			continue
		}

		// TODO: This still doesn't delete the latest, test now to see what happens
		latest := image.Latest(images)

		for _, i := range images {
			if strings.HasSuffix(i.Path, constant.LatestSuffix) {
				continue
			}

			if i.Path == latest.Path {
				continue
			}

			fmt.Printf("Image: %s\n", i.Name)
			c.DeleteImage(p.Identifier, r.ID, i.Name)
		}
	}
}
