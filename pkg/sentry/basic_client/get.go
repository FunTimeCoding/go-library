package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/strings/join/key_value"
	"github.com/funtimecoding/go-library/pkg/web"
)

func (c *Client) Get(path string) string {
	r := web.NewGet("https://%s/api/0%s", c.host, path)
	r.Header.Add(web.ContentTypeHeader, web.ObjectContentType)
	r.Header.Add(
		web.AuthorizationHeader,
		key_value.Space(web.BearerPrefix, c.token),
	)
	r.Header.Add(web.AcceptHeader, web.ObjectContentType)
	response := web.Send(web.Client(true), r)
	fmt.Println(r)
	fmt.Println(response)

	return web.ReadString(response)
}
