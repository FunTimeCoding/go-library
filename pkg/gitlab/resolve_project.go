package gitlab

import (
	"fmt"
	"strconv"
)

func (c *Client) ResolveProject(identifier string) (int64, error) {
	id, e := strconv.ParseInt(identifier, 10, 64)

	if e == nil {
		return id, nil
	}

	result, _, e := c.client.Projects.GetProject(identifier, nil)

	if e != nil {
		return 0, fmt.Errorf("project %s: %w", identifier, e)
	}

	return int64(result.ID), nil
}
