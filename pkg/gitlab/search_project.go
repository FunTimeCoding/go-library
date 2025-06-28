package gitlab

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) SearchProject(query string) []*gitlab.Project {
	result, _, e := c.client.Search.Projects(query, &gitlab.SearchOptions{})
	errors.PanicOnError(e)

	return result
}
