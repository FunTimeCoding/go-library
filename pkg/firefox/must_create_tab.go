package firefox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/firefox/tab"
)

func (c *Client) MustCreateTab(l string) *tab.Tab {
	result, e := c.CreateTab(l)
	errors.PanicOnError(e)

	return result
}
