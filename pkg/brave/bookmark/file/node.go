package file

type Node struct {
	DateAdded    string  `json:"date_added"`
	DateLastUsed string  `json:"date_last_used"`
	Identifier   string  `json:"guid"`
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Locator      string  `json:"url"`
	Children     []*Node `json:"children"`
}
