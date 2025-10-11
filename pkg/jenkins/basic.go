package jenkins

import "github.com/funtimecoding/go-library/pkg/jenkins/basic"

func (c *Client) Basic() *basic.Client {
	return c.basic
}
