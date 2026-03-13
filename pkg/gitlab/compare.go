package gitlab

import "gitlab.com/gitlab-org/api/client-go/v2"

func (c *Client) Compare(
	project int64,
	from, to string,
) *gitlab.Compare {
	result, r, e := c.client.Repositories.Compare(
		project,
		&gitlab.CompareOptions{From: new(from), To: new(to)},
	)
	panicOnError(r, e)

	return result
}
