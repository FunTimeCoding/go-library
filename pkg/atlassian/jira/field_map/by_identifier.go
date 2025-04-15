package field_map

import "github.com/andygrunwald/go-jira"

func (m *Map) ByIdentifier(s string) *jira.Field {
	for _, f := range m.fields {
		if f.ID == s {
			return f
		}
	}

	return nil
}
