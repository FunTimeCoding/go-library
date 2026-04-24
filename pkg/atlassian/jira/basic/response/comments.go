package response

import "github.com/andygrunwald/go-jira"

type Comments struct {
	StartAt    int             `json:"startAt"`
	MaxResults int             `json:"maxResults"`
	Total      int             `json:"total"`
	Comments   []*jira.Comment `json:"comments"`
}
