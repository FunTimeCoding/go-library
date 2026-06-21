package mock_client

import (
	"fmt"
	"gitlab.com/gitlab-org/api/client-go/v2"
)

func (c *Client) File(
	_ int64,
	_ string,
	name string,
) (*gitlab.File, error) {
	f, exists := c.files[name]

	if !exists {
		return nil, fmt.Errorf("file not found: %s", name)
	}

	return f, nil
}
