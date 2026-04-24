package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) allComments(key string) []*jira.Comment {
	var result []*jira.Comment
	startAt := 0

	for {
		p := c.commentPage(key, startAt)
		result = append(result, p.Comments...)

		if startAt+p.MaxResults >= p.Total {
			break
		}

		startAt += p.MaxResults
	}

	return result
}
