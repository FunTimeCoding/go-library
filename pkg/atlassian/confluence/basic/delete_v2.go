package basic

import "github.com/funtimecoding/go-library/pkg/web"

func (c *Client) DeleteV2(l string) string {
	r := web.NewDelete(l)
	r.SetBasicAuth(c.user, c.token)

	return web.ReadString(web.Send(web.Client(), r))
}
