package response

type JiraUser struct {
	Self              string        `json:"self"`
	AccountIdentifier string        `json:"accountId"`
	EmailAddress      string        `json:"emailAddress"`
	AvatarLocator     AvatarLocator `json:"avatarUrls"`
	DisplayName       string        `json:"displayName"`
	Active            bool          `json:"active"`
	TimeZone          string        `json:"timeZone"`
	AccountType       string        `json:"accountType"`
}
