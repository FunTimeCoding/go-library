package field_map

import "github.com/andygrunwald/go-jira"

func (m *Map) ByIdentifier(s string) *jira.Field {
	for _, element := range m.fields {
		if element.ID == s {
			return element
		}
	}

	return nil
}
