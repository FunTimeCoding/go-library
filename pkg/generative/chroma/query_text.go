package chroma

import "github.com/amikos-tech/chroma-go/pkg/api/v2"

func (c *Client) QueryText(
	l v2.Collection,
	t string,
) v2.QueryResult {
	return c.Query(l, v2.WithQueryTexts(t))
}
