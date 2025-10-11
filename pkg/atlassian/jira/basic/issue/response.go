package issue

type Values struct {
	Expands       []string `json:"_expands"`
	IssueId       string   `json:"issueId"`
	IssueKey      string   `json:"issueKey"`
	Summary       string   `json:"summary"`
	RequestTypeId string   `json:"requestTypeId"`
	ServiceDeskId string   `json:"serviceDeskId"`
	CreatedDate   struct {
		Iso8601     string `json:"iso8601"`
		Jira        string `json:"jira"`
		Friendly    string `json:"friendly"`
		EpochMillis int64  `json:"epochMillis"`
	} `json:"createdDate"`
	Reporter struct {
		AccountId    string `json:"accountId"`
		EmailAddress string `json:"emailAddress"`
		DisplayName  string `json:"displayName"`
		Active       bool   `json:"active"`
		TimeZone     string `json:"timeZone"`
		Links        struct {
			JiraRest   string `json:"jiraRest"`
			AvatarUrls struct {
				X48 string `json:"48x48"`
				X24 string `json:"24x24"`
				X16 string `json:"16x16"`
				X32 string `json:"32x32"`
			} `json:"avatarUrls"`
			Self string `json:"self"`
		} `json:"_links"`
	} `json:"reporter"`
	RequestFieldValues []RequestFieldValue `json:"requestFieldValues"`
	CurrentStatus      struct {
		Status         string `json:"status"`
		StatusCategory string `json:"statusCategory"`
		StatusDate     struct {
			Iso8601     string `json:"iso8601"`
			Jira        string `json:"jira"`
			Friendly    string `json:"friendly"`
			EpochMillis int64  `json:"epochMillis"`
		} `json:"statusDate"`
	} `json:"currentStatus"`
	Links struct {
		Web   string `json:"web"`
		Agent string `json:"agent"`
		Self  string `json:"self"`
	} `json:"_links"`
}

type RequestFieldValue struct {
	FieldId       string `json:"fieldId"`
	Label         string `json:"label"`
	Value         any    `json:"value"`
	RenderedValue struct {
		Html string `json:"html"`
	} `json:"renderedValue,omitempty"`
}

type Response struct {
	Expands    []string `json:"_expands"`
	Size       int      `json:"size"`
	Start      int      `json:"start"`
	Limit      int      `json:"limit"`
	IsLastPage bool     `json:"isLastPage"`
	Links      struct {
		Self    string `json:"self"`
		Base    string `json:"base"`
		Context string `json:"context"`
		Next    string `json:"next"`
	} `json:"_links"`
	Values []Values `json:"values"`
}
