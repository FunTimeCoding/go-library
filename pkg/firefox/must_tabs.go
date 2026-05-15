package firefox

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/firefox/tab"
)

func (c *Client) MustTabs() []tab.Tab {
	result, e := c.Tabs()
	errors.PanicOnError(e)

	return result
}
