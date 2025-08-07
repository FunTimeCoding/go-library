package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
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
	r.Header[constant.AcceptHeader] = []string{constant.ObjectContentType}
	r.Header[constant.ContentTypeHeader] = []string{constant.ObjectContentType}

	return web.ReadString(web.Send(web.Client(true), r))
}
