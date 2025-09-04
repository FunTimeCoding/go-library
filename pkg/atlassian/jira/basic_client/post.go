package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
)

func (c *Client) Post(
	path string,
	body string,
) (int, string) {
	r := web.NewPost(fmt.Sprintf("%s%s", c.locator, path), body)
	r.SetBasicAuth(c.user, c.token)
	r.Header.Add(constant.ContentTypeHeader, constant.ObjectContentType)
	r.Header.Add(constant.AcceptHeader, constant.ObjectContentType)
	response := web.Send(web.Client(true), r)

	return response.StatusCode, web.ReadString(response)
}
