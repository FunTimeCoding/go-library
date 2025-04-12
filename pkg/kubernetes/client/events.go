package client

import (
	"github.com/funtimecoding/go-library/pkg/kubernetes/client/operation/list"
	"github.com/funtimecoding/go-library/pkg/kubernetes/types/native/event"
)

func (c *Client) Events(
	namespace string,
	limit int,
	fieldSelector string,
) []*event.Event {
	var result []*event.Event

	for _, client := range c.selectClients(nil) {
		result = append(
			result,
			event.NewSlice(
				list.Event(
					client.client,
					client.context,
					namespace,
					limit,
					fieldSelector,
				),
				client.cluster,
			)...,
		)
	}

	return result
}
