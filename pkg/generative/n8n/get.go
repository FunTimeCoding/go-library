package n8n

import (
	"github.com/funtimecoding/go-library/pkg/generative/n8n/constant"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) Get(path string) string {
	r := web.NewGet("https://%s/api/v1%s", c.host, path)
	r.Header.Add(constant.TokenHeader, c.token)

	return web.ReadString(web.Send(web.Client(true), r))
}
