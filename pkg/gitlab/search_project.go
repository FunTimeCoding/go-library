package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/gitlab/project"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) SearchProject(query string) []*project.Project {
	result, r, e := c.client.Search.Projects(query, &gitlab.SearchOptions{})
	panicOnError(r, e)

	return project.NewSlice(result)
}
