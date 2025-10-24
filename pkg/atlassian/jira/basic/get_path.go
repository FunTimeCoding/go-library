package basic

import (
	"github.com/funtimecoding/go-library/pkg/web"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) GetPath(path string) (int, string) {
	r := web.NewGet(locator.New(c.host).Path(path).String())
	r.SetBasicAuth(c.user, c.token)
	response := web.Send(web.Client(), r)

	return response.StatusCode, web.ReadString(response)
}
