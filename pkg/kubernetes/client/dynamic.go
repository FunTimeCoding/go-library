package client

import "k8s.io/client-go/dynamic"

func (c *Client) Dynamic() dynamic.Interface {
	return c.dynamic
}
