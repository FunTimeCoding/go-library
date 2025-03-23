package team

import "github.com/opsgenie/opsgenie-go-sdk-v2/team"

func NewSlice(v []team.ListedTeams) []*Team {
	var result []*Team

	for _, element := range v {
		result = append(result, New(&element))
	}

	return result
}
