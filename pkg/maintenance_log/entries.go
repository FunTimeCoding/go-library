package maintenance_log

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/client"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) Entries() string {
	result, e := c.client.GetEntries(c.context, &client.GetEntriesParams{})
	errors.PanicOnError(e)

	return web.ReadString(result)
}
