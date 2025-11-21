package request_context

import "github.com/funtimecoding/go-library/pkg/web/constant"

func (c *Context) WriteSuccess() {
	c.WriteOkay(constant.Success)
}
