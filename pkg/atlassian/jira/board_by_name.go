package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) BoardByName(name string) *jira.Board {
	for _, b := range c.Boards() {
		if b.Name == name {
			return b
		}
	}

	return nil
}
