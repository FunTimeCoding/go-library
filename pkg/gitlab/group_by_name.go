package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"gitlab.com/gitlab-org/api/client-go"
)

func (c *Client) GroupByName(s string) *gitlab.Group {
	result, r, e := c.client.Groups.SearchGroup(s)
	panicOnError(r, e)
	fmt.Printf("Groups: %+v\n", result)

	if l := len(result); l != 1 {
		unexpected.Integer(l)
	}

	return result[0]
}
