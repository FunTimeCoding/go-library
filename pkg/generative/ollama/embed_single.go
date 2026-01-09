package ollama

func (c *Client) EmbedSingle(
	model string,
	v string,
) []float32 {
	return c.Embed(model, []string{v})[0]
}
