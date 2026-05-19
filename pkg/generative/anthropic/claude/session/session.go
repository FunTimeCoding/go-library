package session

type Session struct {
	Identifier string
	Slug       string
	Timestamp  string
	CWD        string
	Branch     string
	Lines      int
}
