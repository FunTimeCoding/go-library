package request_context

import "github.com/funtimecoding/go-library/pkg/system"

func (c *Context) Body() string {
	return string(system.ReadAll(c.request.Body))
}
