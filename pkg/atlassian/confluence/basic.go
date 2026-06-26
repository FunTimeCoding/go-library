package confluence

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic"

func (c *Client) Basic() *basic.Client {
	return c.basic
}
