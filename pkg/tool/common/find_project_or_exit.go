package common

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"os"
)

func FindProjectOrExit(
	c *gitlab.Client,
	owner string,
	repository string,
) *project.Project {
	p := c.ProjectByName(owner, repository)

	if p == nil {
		fmt.Printf("repository not found: %s/%s\n", owner, repository)

		os.Exit(1)
	}

	return p
}
