package request_context

func (c *Context) BodyBytes() []byte {
	c.readBody()

	return c.bodyByte
}
