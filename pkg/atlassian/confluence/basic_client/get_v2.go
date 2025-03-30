package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) GetV2(path string) string {
	r := web.NewGet(
		fmt.Sprintf("https://%s/wiki/api/v2%s", c.host, path),
	)
	r.SetBasicAuth(c.user, c.token)

	return web.ReadString(web.Send(web.Client(true), r))
}
