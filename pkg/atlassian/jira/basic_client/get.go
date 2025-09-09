package basic_client

import "github.com/funtimecoding/go-library/pkg/web"

func (c *Client) Get(path string) (int, string) {
	r := web.NewGet("%s%s", c.locator, path)
	r.SetBasicAuth(c.user, c.token)
	response := web.Send(web.Client(true), r)

	return response.StatusCode, web.ReadString(response)
}
