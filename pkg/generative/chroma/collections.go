package chroma

import (
	"github.com/amikos-tech/chroma-go/pkg/api/v2"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) Collections() []v2.Collection {
	result, e := c.client.ListCollections(c.context)
	errors.PanicOnError(e)

	return result
}
