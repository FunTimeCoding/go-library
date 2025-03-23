package field_map

import "github.com/andygrunwald/go-jira"

func (m *Map) ByName(s string) *jira.Field {
	for _, element := range m.fields {
		if element.Name == s {
			return element
		}
	}

	return nil
}
