package jira

import "github.com/funtimecoding/go-library/pkg/atlassian/jira/field_map"

func (c *Client) FieldMap() *field_map.Map {
	if c.fieldMap != nil {
		return c.fieldMap
	}

	c.fieldMap = field_map.New(c.Fields())

	return c.fieldMap
}
