package enriched_session

type Session struct {
	Identifier  string
	Slug        string
	Timestamp   string
	CWD         string
	Branch      string
	Lines       int
	Name        string
	Alias       string
	Description string
	TurnCount   int
	Active      bool
}
