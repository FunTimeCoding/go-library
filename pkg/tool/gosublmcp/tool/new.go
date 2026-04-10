package tool

import "github.com/funtimecoding/go-library/pkg/sublime"

func New(c *sublime.Client) *Tool {
	return &Tool{client: c}
}
