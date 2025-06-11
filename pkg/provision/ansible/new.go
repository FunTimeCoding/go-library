package ansible

import "context"

func New(inventory string) *Client {
	return &Client{context: context.Background(), inventory: inventory}
}
