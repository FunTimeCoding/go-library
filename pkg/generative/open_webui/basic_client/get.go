package basic_client

import "github.com/funtimecoding/go-library/pkg/web"

func (c *Client) Get(path string) string {
	r := web.NewGet("%s://%s%s", web.SecureScheme, c.host, path)
	web.Bearer(r, c.token)

	return web.ReadString(web.Send(web.Client(true), r))
}
