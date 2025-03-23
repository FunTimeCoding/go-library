package field_map

import "github.com/andygrunwald/go-jira"

func New(v []jira.Field) *Map {
	result := &Map{}

	for _, element := range v {
		result.fields = append(result.fields, &element)
	}

	return result
}
