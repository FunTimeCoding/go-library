package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"github.com/xanzy/go-gitlab"
	"log"
)

func (c *Client) ProjectByName(
	namespace string,
	name string,
) *project.Project {
	result, _, e := c.client.Projects.ListProjects(
		&gitlab.ListProjectsOptions{Search: &name},
	)
	errors.PanicOnError(e)
	count := len(result)

	if count == 0 {
		return nil
	} else if count > 1 {
		log.Panicf("unexpected: %d", count)
	}

	r := result[0]

	if r.Namespace.Path == namespace {
		return project.New(r)
	}

	return nil
}
