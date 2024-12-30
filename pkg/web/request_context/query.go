package request_context

func (c *Context) Query(k string) string {
	return c.Request().URL.Query().Get(k)
}
