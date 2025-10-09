package basic_client

import "github.com/funtimecoding/go-library/pkg/web"

func (c *Client) Get(path string) string {
	r := web.NewGet("%s://%s/wiki/rest/api%s", c.scheme, c.host, path)
	r.SetBasicAuth(c.user, c.token)

	return web.ReadString(web.Send(web.Client(true), r))
}
