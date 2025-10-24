package basic

import "github.com/funtimecoding/go-library/pkg/web"

func (c *Client) Get(l string) string {
	r := web.NewGet(l)
	r.SetBasicAuth(c.user, c.token)

	return web.ReadString(web.Send(web.Client(), r))
}
