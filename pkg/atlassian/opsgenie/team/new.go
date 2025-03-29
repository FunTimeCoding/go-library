package team

import "github.com/opsgenie/opsgenie-go-sdk-v2/team"

func New(v *team.ListedTeams) *Team {
	return &Team{
		Identifier:  v.Id,
		Name:        v.Name,
		Description: v.Description,
		Raw:         v,
	}
}
