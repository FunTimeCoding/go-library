package model_context

func displayName(
	c *caller,
	target string,
) string {
	if target != "" {
		return target
	}

	return c.Callsign
}
