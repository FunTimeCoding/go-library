package sublime

import "github.com/funtimecoding/go-library/pkg/web/locator"

func WithHost(host string) Option {
	return func(c *Client) { c.base = locator.New(host).Insecure() }
}
