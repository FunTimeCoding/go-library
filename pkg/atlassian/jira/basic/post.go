package basic

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/constant"
)

func (c *Client) Post(
	l string,
	body string,
) (int, string) {
	r := web.NewPost(l, body)
	r.SetBasicAuth(c.user, c.token)
	r.Header.Add(constant.ContentType, constant.Object)
	r.Header.Add(constant.Accept, constant.Object)
	response := web.Send(web.Client(), r)

	return response.StatusCode, web.ReadString(response)
}
