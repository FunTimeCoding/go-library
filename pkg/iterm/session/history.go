package session

type History struct {
	Identifier      string   `json:"session_id"`
	Lines           []string `json:"lines"`
	StartLine       int      `json:"start_line"`
	TotalScrollback int      `json:"total_scrollback"`
}
