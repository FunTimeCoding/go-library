package salt

import "github.com/funtimecoding/go-library/pkg/provision/salt/basic/response"

func (c *Client) Jobs() ([]response.Job, error) {
	return c.basic.ListJobs()
}
