package telemetry

func (c *Client) Record(
	tool string,
	surface string,
	actor string,
	outcome string,
) {
	if c == nil {
		return
	}

	go c.send(tool, surface, actor, outcome, 0)
}
