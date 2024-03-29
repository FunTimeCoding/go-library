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
	count := len(result)

	if count != 1 {
		log.Panicf("unexpected: %d", count)
	}

	return result[0]
}
