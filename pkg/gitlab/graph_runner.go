package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/gitlab/response"
)

func (c *Client) GraphRunner(identifier int) *response.Runner {
	result := &response.Runner{}
	c.Query(
		fmt.Sprintf(
			"query {runner(id: \"gid://gitlab/Ci::Runner/%d\") { id description status runnerType managers { nodes { systemId ipAddress version revision } } } }",
			identifier,
		),
		&result,
	)

	return result
}
