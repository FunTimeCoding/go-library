package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) PostV2(
	path string,
	body string,
) string {
	r := web.NewPost(
		fmt.Sprintf("https://%s/wiki/api/v2%s", c.host, path),
		body,
	)
	r.SetBasicAuth(c.user, c.token)
	r.Header[web.AcceptHeader] = []string{web.ObjectContentType}
	r.Header[web.ContentTypeHeader] = []string{web.ObjectContentType}

	return web.ReadString(web.Send(web.Client(true), r))
}
