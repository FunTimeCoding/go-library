package tab

type Tab struct {
	Identifier int    `json:"tab_id"`
	WindowId   int    `json:"window_id"`
	Locator    string `json:"url"`
	Title      string `json:"title"`
	Active     bool   `json:"active"`
	Status     string `json:"status"`
	Index      int    `json:"index"`
	GroupId    int    `json:"group_id"`
}
