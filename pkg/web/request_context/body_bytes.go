package request_context

func (c *Context) BodyByte() []byte {
	c.readBody()

	return c.bodyByte
}
