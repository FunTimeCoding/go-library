package salt

func (c *Client) LocalAsync(
	target string,
	function string,
	a []string,
) (string, error) {
	return c.basic.LocalClientAsync(target, function, a)
}
