package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
)

func (c *Client) Get(path string) string {
	r := web.NewGet("https://%s/api/0%s", c.host, path)
	r.Header.Add(constant.ContentTypeHeader, constant.ObjectContentType)
	r.Header.Add(
		constant.AuthorizationHeader,
		key_value.Space(constant.BearerPrefix, c.token),
	)
	r.Header.Add(constant.AcceptHeader, constant.ObjectContentType)
	response := web.Send(web.Client(true), r)
	fmt.Println(r)
	fmt.Println(response)

	return web.ReadString(response)
}
