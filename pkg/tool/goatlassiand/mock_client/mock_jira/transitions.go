package mock_jira

import "github.com/andygrunwald/go-jira"

func (c *Client) Transitions(string) ([]jira.Transition, error) { return nil, nil }
