package confluence

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/confluence/search_result"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/web/locator"
)

func (c *Client) Search(
	query string,
	a ...any,
) ([]*search_result.Result, error) {
	if len(a) > 0 {
		query = fmt.Sprintf(query, a...)
	}

	body, e := c.basic.Get(
		locator.New(c.host).Base(constant.OldBase).Path(
			constant.Search,
		).Set(constant.Query, query).String(),
	)

	if e != nil {
		return nil, e
	}

	var result *response.Search
	notation.DecodeStrict(body, &result, false)

	return search_result.NewSlice(result.Results), nil
}
