package mock_jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Projects() (*jira.ProjectList, error) { return nil, nil }
