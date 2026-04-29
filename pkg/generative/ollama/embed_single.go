package ollama

func (c *Client) EmbedSingle(
	model string,
	v string,
) ([]float32, error) {
	result, e := c.Embed(model, []string{v})

	if e != nil {
		return nil, e
	}

	return result[0], nil
}
