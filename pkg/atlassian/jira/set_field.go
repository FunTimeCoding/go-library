package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) SetField(
	i *jira.Issue,
	name string,
	value any,
) {
	i.Fields.Unknowns.Set(c.FieldMap().ByName(name).Key, value)
}
