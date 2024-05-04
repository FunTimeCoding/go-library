package request_context

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Context) String(code int, format string, values ...any) {
	c.writer.WriteHeader(code)
	_, e := c.writer.Write([]byte(fmt.Sprintf(format, values...)))
	errors.PanicOnError(e)
}
