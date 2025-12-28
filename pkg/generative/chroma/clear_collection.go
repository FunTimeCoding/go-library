package chroma

import "github.com/amikos-tech/chroma-go/pkg/api/v2"

func (c *Client) ClearCollection(l v2.Collection) {
	// Deletes all, even if no name field exists
	c.Delete(l, v2.WithWhereDelete(v2.NotEqString("name", "")))
}
