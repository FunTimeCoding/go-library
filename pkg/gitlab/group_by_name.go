package gitlab

import (
	"fmt"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) GroupByName(s string) (*gitlab.Group, error) {
	result, _, e := c.client.Groups.SearchGroup(s)

	if e != nil {
		return nil, e
	}

	if len(result) != 1 {
		return nil, fmt.Errorf("expected 1 group, got %d", len(result))
	}

	return result[0], nil
}
