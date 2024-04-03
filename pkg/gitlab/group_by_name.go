package gitlab

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/xanzy/go-gitlab"
	"log"
)

func (c *Client) GroupByName(s string) *gitlab.Group {
	result, _, e := c.client.Groups.SearchGroup(s)
	errors.PanicOnError(e)
	fmt.Printf("Groups: %+v\n", result)

	if l := len(result); l != 1 {
		log.Panicf("unexpected: %d", l)
	}

	return result[0]
}
