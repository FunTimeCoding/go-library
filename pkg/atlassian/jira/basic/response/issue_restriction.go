package response

type IssueRestriction struct {
	IssueRestrictions struct{} `json:"issuerestrictions"`
	ShouldDisplay     bool     `json:"shouldDisplay"`
}
