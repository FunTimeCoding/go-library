package mock_client

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/search_result"

func (c *Client) Search(
	query string,
	a ...any,
) ([]*search_result.Result, error) {
	return nil, nil
}
