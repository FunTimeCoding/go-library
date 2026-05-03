package issue

type User struct {
	AccountId    string    `json:"accountId"`
	EmailAddress string    `json:"emailAddress"`
	DisplayName  string    `json:"displayName"`
	Active       bool      `json:"active"`
	TimeZone     string    `json:"timeZone"`
	Links        UserLinks `json:"_links"`
}
