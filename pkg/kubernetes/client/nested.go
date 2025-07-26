package client

import "k8s.io/client-go/kubernetes"

func (c *Client) Nested() *kubernetes.Clientset {
	return c.client
}
