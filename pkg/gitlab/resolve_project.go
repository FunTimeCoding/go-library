package gitlab

import (
	"fmt"
	"strconv"
)

func (c *Client) ResolveProject(identifier string) (int64, error) {
	i, e := strconv.ParseInt(identifier, 10, 64)

	if e == nil {
		return i, nil
	}

	result, _, f := c.client.Projects.GetProject(identifier, nil)

	if f != nil {
		return 0, fmt.Errorf("project %s: %w", identifier, f)
	}

	return result.ID, nil
}
