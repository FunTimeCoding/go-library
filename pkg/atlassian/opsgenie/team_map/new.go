package team_map

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/team"

func New(v []*team.Team) *Map {
	m := make(map[string]*team.Team, len(v))

	for _, element := range v {
		m[element.Identifier] = element
	}

	result := &Map{
		Teams:        v,
		KeyByNameMap: make(map[string]string),
		TeamMap:      m,
	}

	return result
}
