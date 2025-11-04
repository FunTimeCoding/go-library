package gitlab

import "gitlab.com/gitlab-org/api/client-go"

func (c *Client) Query(
	q string,
	response any,
) {
	r, e := c.client.GraphQL.Do(
		gitlab.GraphQLQuery{Query: q},
		response,
		gitlab.WithContext(c.context),
	)
	panicOnError(r, e)
}
