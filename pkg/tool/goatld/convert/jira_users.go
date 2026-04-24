package convert

import "github.com/andygrunwald/go-jira"

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
