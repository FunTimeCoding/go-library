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
) ([]custom_field_value.Value, error) {
	m, e := c.FieldMap()

	if e != nil {
		return nil, e
	}

	f := m.ByName(fieldName)

	if f == nil {
		return nil, nil
	}

	meta, g := c.CreateMeta(key)

	if g != nil {
		return nil, g
	}

	for _, p := range meta.Projects {
		for _, t := range p.IssueTypes {
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

			return result, nil
		}
	}

	return nil, nil
}
