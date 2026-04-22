package convert

import "github.com/andygrunwald/go-jira"

type JiraUser struct {
	AccountIdentifier string `json:"account_identifier"`
	DisplayName       string `json:"display_name"`
	Email             string `json:"email,omitempty"`
	Active            bool   `json:"active"`
}

func JiraUsers(users []jira.User) []*JiraUser {
	var result []*JiraUser

	for _, u := range users {
		result = append(
			result,
			&JiraUser{
				AccountIdentifier: u.AccountID,
				DisplayName:       u.DisplayName,
				Email:             u.EmailAddress,
				Active:            u.Active,
			},
		)
	}

	return result
}
