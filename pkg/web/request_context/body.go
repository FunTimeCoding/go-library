package request_context

func (c *Context) Body() string {
	c.readBody()

	return c.body
}
