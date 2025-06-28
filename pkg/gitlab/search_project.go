package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) SearchProject(query string) []*project.Project {
	result, _, e := c.client.Search.Projects(query, &gitlab.SearchOptions{})
	errors.PanicOnError(e)

	return project.NewSlice(result)
}
