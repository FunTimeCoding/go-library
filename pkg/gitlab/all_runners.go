package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) allRunners(o *gitlab.ListRunnersOptions) ([]*gitlab.Runner, *gitlab.Response) {
	result, r, e := c.client.Runners.ListAllRunners(o)
	panicOnError(r, e)

	return result, r
}
