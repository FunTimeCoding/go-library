package gitlab

import "github.com/funtimecoding/go-library/pkg/gitlab/merge_request"

func (c *Client) ProjectsMergeRequests() []*merge_request.Request {
	var result []*merge_request.Request

	for _, identifier := range c.projects {
		for _, element := range c.ProjectMergeRequests(identifier) {
			if element.Done() {
				continue
			}

			result = append(result, element)
		}
	}

	return result
}
