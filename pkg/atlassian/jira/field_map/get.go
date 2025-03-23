package field_map

import "github.com/andygrunwald/go-jira"

func (m *Map) Get() []*jira.Field {
	return m.fields
}
