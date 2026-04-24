package jira

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/issue"
)

func (c *Client) FillComments(v []*issue.Issue) {
	filled := 0

	for _, i := range v {
		if i.Raw.Fields.Comments == nil {
			continue
		}

		if len(i.Raw.Fields.Comments.Comments) < constant.CommentCap {
			continue
		}

		all := c.allComments(i.Key)
		i.Raw.Fields.Comments.Comments = all
		filled++
	}

	if filled > 0 {
		fmt.Printf(
			"Filled comments for %d issue(s)\n",
			filled,
		)
	}
}
