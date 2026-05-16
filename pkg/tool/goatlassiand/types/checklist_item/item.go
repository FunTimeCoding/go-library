package checklist_item

type Item struct {
	Index   int    `json:"index"`
	Text    string `json:"text"`
	Checked bool   `json:"checked"`
}
