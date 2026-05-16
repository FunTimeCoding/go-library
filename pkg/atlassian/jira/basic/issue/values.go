package issue

type Values struct {
	Expands               []string            `json:"_expands"`
	IssueIdentifier       string              `json:"issueId"`
	IssueKey              string              `json:"issueKey"`
	Summary               string              `json:"summary"`
	RequestTypeIdentifier string              `json:"requestTypeId"`
	ServiceDeskIdentifier string              `json:"serviceDeskId"`
	CreatedDate           JiraDate            `json:"createdDate"`
	Reporter              User                `json:"reporter"`
	RequestFieldValues    []RequestFieldValue `json:"requestFieldValues"`
	CurrentStatus         CurrentStatus       `json:"currentStatus"`
	Links                 ValueLinks          `json:"_links"`
}
