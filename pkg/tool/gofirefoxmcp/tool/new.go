package tool

import "github.com/funtimecoding/go-library/pkg/firefox"

func New(c *firefox.Client) *Tool {
	return &Tool{client: c}
}
