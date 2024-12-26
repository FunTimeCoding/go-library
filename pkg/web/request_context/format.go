package request_context

import "fmt"

func (c *Context) Format(
	code int,
	format string,
	values ...any,
) {
	c.Write(code, fmt.Sprintf(format, values...))
}
