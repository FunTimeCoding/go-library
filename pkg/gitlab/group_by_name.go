package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/unexpected"
	"github.com/xanzy/go-gitlab"
)

func (c *Client) GroupByName(s string) *gitlab.Group {
	result, _, e := c.client.Groups.SearchGroup(s)
	errors.PanicOnError(e)
	fmt.Printf("Groups: %+v\n", result)

	if l := len(result); l != 1 {
		unexpected.Integer(l)
	}

	return result[0]
}
