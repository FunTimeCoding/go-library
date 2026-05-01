package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/field_map"
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (c *Client) MustFieldMap() *field_map.Map {
	result, e := c.FieldMap()
	errors.PanicOnError(e)

	return result
}
