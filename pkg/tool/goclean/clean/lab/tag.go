package lab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"github.com/funtimecoding/go-library/pkg/gitlab/tag"
)

func Tag(
	c *gitlab.Client,
	p *project.Project,
) {
	tags := c.Tags(p.Identifier)

	if len(tags) == 0 {
		return
	}

	latest := tag.Latest(tags)

	for _, t := range tags {
		if t.Name == latest.Name {
			continue
		}

		fmt.Printf("Delete tag: %s\n", t.Name)
		c.DeleteTag(p.Identifier, t.Name)
	}

	git.Fetch()
}
