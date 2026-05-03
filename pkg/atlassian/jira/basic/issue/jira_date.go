package issue

type JiraDate struct {
	Iso8601     string `json:"iso8601"`
	Jira        string `json:"jira"`
	Friendly    string `json:"friendly"`
	EpochMillis int64  `json:"epochMillis"`
}
