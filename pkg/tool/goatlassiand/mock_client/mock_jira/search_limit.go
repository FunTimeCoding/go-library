package mock_jira

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"

func (c *Client) SearchLimit(
	int,
	string,
	...any,
) ([]*issue.Issue, error) {
	return nil, nil
}
