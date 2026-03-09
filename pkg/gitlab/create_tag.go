package gitlab

import "gitlab.com/gitlab-org/api/client-go"

func (c *Client) CreateTag(
	project int64,
	name string,
	reference string,
	message string,
) *gitlab.Tag {
	result, r, e := c.client.Tags.CreateTag(
		project,
		&gitlab.CreateTagOptions{
			TagName: new(name),
			Ref:     new(reference),
			Message: new(message),
		},
	)
	panicOnError(r, e)

	return result
}
