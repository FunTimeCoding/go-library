package field_map

import "github.com/andygrunwald/go-jira"

func (m *Map) ByName(s string) *jira.Field {
	for _, f := range m.fields {
		if f.Name == s {
			return f
		}
	}

	return nil
}
