package notification

type Notification struct {
	User     string `json:"user"`
	Token    string `json:"token"`
	Message  string `json:"message"`
	Priority int    `json:"priority"`
}
