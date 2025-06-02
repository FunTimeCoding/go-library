package chroma

import "github.com/amikos-tech/chroma-go/pkg/api/v2"

func (c *Client) DeleteIdentifiers(
	l v2.Collection,
	i ...v2.DocumentID,
) {
	c.Delete(l, v2.WithIDsDelete(i...))
}
