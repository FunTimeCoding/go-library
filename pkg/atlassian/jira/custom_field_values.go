package jira

import (
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/custom_field_value"
	"github.com/funtimecoding/go-library/pkg/notation"
)

func (c *Client) CustomFieldValues(
	key string,
	issueType string,
	fieldName string,
) []custom_field_value.Value {
	f := c.FieldMap().ByName(fieldName)

	for _, m := range c.CreateMeta(key).Projects {
		for _, t := range m.IssueTypes {
			if t.Name != issueType {
				continue
			}

			var result []custom_field_value.Value
			notation.DecodeStrict(
				notation.Encode(
					t.Fields[f.ID].(map[string]any)[constant.AllowedValuesKey],
					false,
				),
				&result,
				true,
			)

			return result
		}
	}

	return nil
}
