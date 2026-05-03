package issue

type UserLinks struct {
	JiraRest   string     `json:"jiraRest"`
	AvatarUrls AvatarUrls `json:"avatarUrls"`
	Self       string     `json:"self"`
}
