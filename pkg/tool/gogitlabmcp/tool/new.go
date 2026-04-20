package tool

import "gitlab.com/gitlab-org/api/client-go/v2"

func New(c *gitlab.Client) *Tool {
	return &Tool{client: c}
}
