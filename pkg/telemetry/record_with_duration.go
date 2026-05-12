package telemetry

func (c *Client) RecordWithDuration(
	tool string,
	surface string,
	actor string,
	outcome string,
	durationMillisecond int64,
) {
	if c == nil {
		return
	}

	go c.send(tool, surface, actor, outcome, durationMillisecond)
}
