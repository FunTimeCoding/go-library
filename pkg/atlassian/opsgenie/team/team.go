package team

import "github.com/opsgenie/opsgenie-go-sdk-v2/team"

type Team struct {
	Identifier  string
	Name        string
	Description string

	Raw *team.ListedTeams
}
