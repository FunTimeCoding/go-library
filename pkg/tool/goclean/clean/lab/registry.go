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
		// TODO: This still doesn't delete the latest, test now to see what happens
		images := c.Images(p.Identifier, r.ID)

		if len(images) == 0 {
			continue
		}

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
