package request_context

import "net/http"

func (c *Context) FormatOkay(format string, values ...any) {
	c.Format(http.StatusOK, format, values...)
}
