package jenkins

import "github.com/funtimecoding/go-library/pkg/jenkins/basic_client"

func (c *Client) Basic() *basic_client.Client {
	return c.basic
}
