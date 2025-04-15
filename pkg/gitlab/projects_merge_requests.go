package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/merge_request"

func (c *Client) ProjectsMergeRequests() []*merge_request.Request {
	var result []*merge_request.Request

	for _, identifier := range c.projects {
		for _, e := range c.ProjectMergeRequests(identifier, false) {
			if e.Done() {
				continue
			}

			result = append(result, e)
		}
	}

	return result
}
