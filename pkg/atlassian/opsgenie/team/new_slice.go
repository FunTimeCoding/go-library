package team

import "github.com/opsgenie/opsgenie-go-sdk-v2/team"

func NewSlice(v []team.ListedTeams) []*Team {
	var result []*Team

	for _, e := range v {
		result = append(result, New(&e))
	}

	return result
}
