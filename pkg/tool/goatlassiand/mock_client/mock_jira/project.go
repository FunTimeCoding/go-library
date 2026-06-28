package mock_jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Project(string) (*jira.Project, error) { return nil, nil }
