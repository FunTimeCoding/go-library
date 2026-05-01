package jira

import "github.com/andygrunwald/go-jira"

func (c *Client) SetField(
	i *jira.Issue,
	name string,
	value any,
) error {
	m, e := c.FieldMap()

	if e != nil {
		return e
	}

	i.Fields.Unknowns.Set(m.ByName(name).Key, value)

	return nil
}
