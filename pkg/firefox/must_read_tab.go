package firefox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/firefox/tab"
)

func (c *Client) MustReadTab(identifier int, raw bool) tab.Content {
	result, e := c.ReadTab(identifier, raw)
	errors.PanicOnError(e)

	return result
}
