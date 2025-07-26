package client

import "k8s.io/client-go/rest"

func (c *Client) Configuration() *rest.Config {
	return c.configuration
}
