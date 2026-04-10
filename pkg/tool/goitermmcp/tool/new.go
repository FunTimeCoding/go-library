package tool

import "github.com/funtimecoding/go-library/pkg/iterm"

func New(c *iterm.Client) *Tool {
	return &Tool{client: c}
}
