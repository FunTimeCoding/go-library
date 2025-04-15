package telegram

import "github.com/funtimecoding/go-library/pkg/bolt"

func (c *Client) SetDatabase(b *bolt.Client) {
	c.database = b
}
