package salt

import "github.com/daixijun/go-salt/v2"

func (c *Client) Jobs() ([]salt.Job, error) {
	return c.client.ListJobs(c.context)
}
