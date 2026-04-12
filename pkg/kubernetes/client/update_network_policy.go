package client

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/operation/update_operation"
	networking "k8s.io/api/networking/v1"
)

func (c *Client) UpdateNetworkPolicy(np *networking.NetworkPolicy) {
	update_operation.NetworkPolicy(c.client, c.context, np)
}
