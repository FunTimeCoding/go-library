package basic

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) Get(path string) string {
	r := web.NewGet(locator.New(c.host).Path(path).String())
	web.Bearer(r, c.token)

	return web.ReadString(web.Send(web.Client(), r))
}
