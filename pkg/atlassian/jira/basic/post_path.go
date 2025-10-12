package basic

import "github.com/funtimecoding/go-library/pkg/web/locator"

func (c *Client) PostPath(
	path string,
	body string,
) (int, string) {
	return c.Post(locator.New(c.host).Path(path).String(), body)
}
