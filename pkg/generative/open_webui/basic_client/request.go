package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) Request(path string) string {
	r := web.NewGet(
		fmt.Sprintf("%s://%s%s", web.SecureScheme, c.host, path),
	)
	web.Bearer(r, c.token)

	return web.ReadString(web.Send(web.Client(true), r))
}
