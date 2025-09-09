package hub

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/git"
	"github.com/funtimecoding/go-library/pkg/github"
	"github.com/funtimecoding/go-library/pkg/github/tag"
	"github.com/funtimecoding/go-library/pkg/system"
)

func Tag(
	c *github.Client,
	namespace string,
	repository string,
) {
	tags := c.Tags(namespace, repository)
	latestTag := tag.Latest(tags)

	if latestTag == nil {
		system.Exitf(
			1,
			"could not find latest tag for %s/%s\n",
			namespace,
			repository,
		)

		return
	}

	for _, t := range tags {
		if t.Name == latestTag.Name {
			continue
		}

		fmt.Printf("Delete tag: %s\n", *t.Name)
		c.DeleteTag(namespace, repository, *t.Name)
	}

	git.Fetch()
}
