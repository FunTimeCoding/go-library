package client

func (c *Client) JobRunning(
	namespace string,
	name string,
) bool {
	j := c.Job(namespace, name)

	if j == nil {
		return false
	}

	return j.Raw.Status.CompletionTime == nil && j.Raw.Status.Failed == 0
}
