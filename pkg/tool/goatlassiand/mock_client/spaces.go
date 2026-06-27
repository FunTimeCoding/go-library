package mock_client

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/space"

func (c *Client) Spaces() ([]*space.Space, error) {
	return nil, nil
}
