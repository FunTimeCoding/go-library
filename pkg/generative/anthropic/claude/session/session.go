package session

type Session struct {
	Identifier    string
	Slug          string
	Timestamp     string
	WorkDirectory string
	Branch        string
	Lines         int
}
