package client

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/operation/get"
	"k8s.io/api/networking/v1"
)

func (c *Client) NetworkPolicy(namespace, name string) *v1.NetworkPolicy {
	return get.NetworkPolicy(c.client, c.context, namespace, name)
}
