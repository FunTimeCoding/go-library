package screen

type Screen struct {
	Identifier string   `json:"session_id"`
	Lines      []string `json:"lines"`
	CursorX    int      `json:"cursor_x"`
	CursorY    int      `json:"cursor_y"`
}
