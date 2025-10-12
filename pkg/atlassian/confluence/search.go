package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/search_result"
	"github.com/funtimecoding/go-library/pkg/notation"
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
			c.basic.Base().Copy().Path(
				constant.Search,
			).Set(constant.QueryKey, query).String(),
		),
		&result,
		false,
	)

	return search_result.NewSlice(result.Results)
}
