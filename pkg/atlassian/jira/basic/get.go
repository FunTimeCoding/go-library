package basic

import "github.com/funtimecoding/go-library/pkg/web"

func (c *Client) Get(l string) (int, string) {
	r := web.NewGet(l)
	r.SetBasicAuth(c.user, c.token)
	response := web.Send(web.Client(), r)

	return response.StatusCode, web.ReadString(response)
}
