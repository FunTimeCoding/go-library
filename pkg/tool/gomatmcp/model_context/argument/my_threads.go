package argument

type MyThreads struct {
	Limit      int    `json:"limit"`
	UnreadOnly bool   `json:"unread_only"`
	Since      string `json:"since"`
}
