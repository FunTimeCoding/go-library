package client

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/operation/get"
	networking "k8s.io/api/networking/v1"
)

func (c *Client) NetworkPolicy(namespace, name string) *networking.NetworkPolicy {
	return get.NetworkPolicy(c.client, c.context, namespace, name)
}
