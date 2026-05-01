package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) allComments(key string) ([]*jira.Comment, error) {
	var result []*jira.Comment
	startAt := 0

	for {
		p, e := c.commentPage(key, startAt)

		if e != nil {
			return nil, e
		}

		result = append(result, p.Comments...)

		if startAt+p.MaxResults >= p.Total {
			break
		}

		startAt += p.MaxResults
	}

	return result, nil
}
