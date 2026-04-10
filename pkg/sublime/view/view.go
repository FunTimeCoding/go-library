package view

type View struct {
	Identifier int    `json:"view_id"`
	WindowId   int    `json:"window_id"`
	FilePath   string `json:"file_path"`
	Title      string `json:"title"`
	Dirty      bool   `json:"is_dirty"`
	Syntax     string `json:"syntax"`
	Preview    string `json:"preview,omitempty"`
	Text       string `json:"text,omitempty"`
}
