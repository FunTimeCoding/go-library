package response

type User struct {
	Type                   string         `json:"type"`
	AccountId              string         `json:"accountId"`
	AccountType            string         `json:"accountType"`
	Email                  string         `json:"email"`
	PublicName             string         `json:"publicName"`
	TimeZone               string         `json:"timeZone"`
	ProfilePicture         ProfilePicture `json:"profilePicture"`
	DisplayName            string         `json:"displayName"`
	IsExternalCollaborator bool           `json:"isExternalCollaborator"`
	IsGuest                bool           `json:"isGuest"`
	Locale                 string         `json:"locale"`
	AccountStatus          string         `json:"accountStatus"`
	Expandable             UserExpandable `json:"_expandable"`
	Links                  Links          `json:"_links"`
}
