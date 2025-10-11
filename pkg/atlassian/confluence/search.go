package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/search_result"
	"github.com/funtimecoding/go-library/pkg/notation"
	"net/url"
)

func (c *Client) Search(
	query string,
	a ...any,
) []*search_result.Result {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	var result *response.Search
	notation.DecodeStrict(
		c.basic.Get(
			fmt.Sprintf(
				"/content/search?cql=%s",
				url.QueryEscape(query),
			),
		),
		&result,
		false,
	)

	return search_result.NewSlice(result.Results)
}
