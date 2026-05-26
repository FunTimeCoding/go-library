package hetzner

import (
	"context"
	"github.com/hetznercloud/hcloud-go/v2/hcloud"
)

func New(
	token string,
	o ...Option,
) *Client {
	result := &Client{context: context.Background()}

	for _, p := range o {
		p(result)
	}

	result.client = hcloud.NewClient(hcloud.WithToken(token))

	return result
}
