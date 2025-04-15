package field_map

import "github.com/andygrunwald/go-jira"

func New(v []jira.Field) *Map {
	result := &Map{}

	for _, e := range v {
		result.fields = append(result.fields, &e)
	}

	return result
}
