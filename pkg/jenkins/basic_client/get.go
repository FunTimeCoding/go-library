package basic_client

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
)

func (c *Client) Get(path string) string {
	r := web.NewGet("%s%s", c.locator, path)
	r.SetBasicAuth(c.user, c.password)
	r.Header.Add(constant.ContentTypeHeader, constant.ObjectContentType)
	r.Header.Add(constant.AcceptHeader, constant.ObjectContentType)
	response := web.Send(web.Client(true), r)

	if false {
		fmt.Println(r)
		fmt.Println(response)
	}

	return web.ReadString(response)
}
