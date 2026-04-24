package convert

type JiraUser struct {
	AccountIdentifier string `json:"account_identifier"`
	DisplayName       string `json:"display_name"`
	Email             string `json:"email,omitempty"`
	Active            bool   `json:"active"`
}
