package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/search_result"
	"github.com/funtimecoding/go-library/pkg/notation"
	"net/url"
)

func (c *Client) Search(query string) []*search_result.Result {
	var result *response.Search
	notation.DecodeStrict(
		c.basic.Get(
			fmt.Sprintf("/content/search?cql=%s", url.QueryEscape(query)),
		),
		&result,
		false,
	)

	return search_result.NewSlice(result.Results)
}
