package mock_jira

import "github.com/andygrunwald/go-jira"

func (c *Client) MetaProject(string) (*jira.MetaProject, error) { return nil, nil }
