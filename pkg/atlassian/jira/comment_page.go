package jira

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/basic/response"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) commentPage(
	key string,
	startAt int,
) *response.Comments {
	var result response.Comments
	_, r := c.basic.Get(
		c.basic.Base().Copy().Base("/rest/api/2").Path(
			fmt.Sprintf("%s/%s/comment", constant.Issue, key),
		).SetInteger(
			constant.MaximumResultsKey,
			constant.CommentPageSize,
		).SetInteger(
			"startAt",
			startAt,
		).String(),
	)
	notation.DecodeStrict(r, &result, true)

	return &result
}
