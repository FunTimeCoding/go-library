package jira

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) searchV3Page(
	maximumResults int,
	nextPageToken string,
	query string,
) *response.Search {
	var result response.Search
	status, r := c.basic.Get(
		c.basic.Base().Copy().Base(constant.Base).Path(constant.Search).Set(
			constant.FieldsKey,
			constant.AllFields,
		).SetInteger(constant.MaximumResultsKey, maximumResults).Set(
			constant.NextPageTokenKey,
			nextPageToken,
		).Set(constant.QueryKey, query).String(),
	)
	notation.DecodeStrict(r, &result, true)

	if false {
		fmt.Printf("Response: %d %s\n", status, r)
	}

	return &result
}
