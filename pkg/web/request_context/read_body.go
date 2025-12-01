package request_context

import "github.com/funtimecoding/go-library/pkg/system"

func (c *Context) readBody() {
	if c.bodyRead {
		return
	}

	c.bodyByte = system.ReadAll(c.request.Body)
	c.body = string(c.bodyByte)
	c.bodyRead = true
}
