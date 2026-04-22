package convert

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/go-library/pkg/atlassian/jira/custom_field_value"
	"github.com/funtimecoding/go-library/pkg/notation"
	"github.com/funtimecoding/go-library/pkg/tool/goatld/constant"
	"sort"
	"strconv"
)

func JiraCreateMeta(
	t *jira.MetaIssueType,
	expand []string,
) []*CreateMetaField {
	expandSet := make(map[string]bool, len(expand))

	for _, n := range expand {
		expandSet[n] = true
	}

	var result []*CreateMetaField

	for key := range t.Fields {
		f := &CreateMetaField{Key: key}
		name, e := t.Fields.String(key + "/name")

		if e == nil {
			f.Name = name
		}

		required, g := t.Fields.Bool(key + "/required")

		if g == nil {
			f.Required = required
		}

		schema, h := t.Fields.String(key + "/schema/type")

		if h == nil {
			f.Schema = schema
		}

		allowed, i := t.Fields.Array(
			key + "/allowedValues",
		)

		if i == nil {
			for _, a := range allowed {
				m, ok := a.(map[string]any)

				if !ok {
					continue
				}

				var v custom_field_value.Value
				notation.DecodeStrict(
					notation.Encode(m, false),
					&v,
					true,
				)

				if v.Value != "" {
					f.AllowedValues = append(
						f.AllowedValues,
						CreateMetaAllowed{
							Identifier: v.ID,
							Value:      v.Value,
						},
					)
				} else if n, ok := m["name"].(string); ok {
					f.AllowedValues = append(
						f.AllowedValues,
						CreateMetaAllowed{
							Identifier: v.ID,
							Value:      n,
						},
					)
				}
			}
		}

		if len(f.AllowedValues) > constant.AllowedValueLimit &&
			!expandSet[f.Name] && !expandSet[f.Key] {
			sort.Slice(
				f.AllowedValues,
				func(a, b int) bool {
					x, xFail := strconv.Atoi(
						f.AllowedValues[a].Identifier,
					)
					y, yFail := strconv.Atoi(
						f.AllowedValues[b].Identifier,
					)

					if xFail != nil || yFail != nil {
						return f.AllowedValues[a].Identifier >
							f.AllowedValues[b].Identifier
					}

					return x > y
				},
			)
			f.AllowedValues = f.AllowedValues[:constant.AllowedValueLimit]
		}

		result = append(result, f)
	}

	sort.Slice(
		result,
		func(a, b int) bool {
			if result[a].Required != result[b].Required {
				return result[a].Required
			}

			return result[a].Name < result[b].Name
		},
	)

	return result
}
