package file

type Node struct {
	DateAdded        string  `json:"date_added"`
	DateLastUsed     string  `json:"date_last_used"`
	GlobalIdentifier string  `json:"guid"`
	Identifier       string  `json:"id"`
	Name             string  `json:"name"`
	Type             string  `json:"type"`
	Locator          string  `json:"url"`
	Children         []*Node `json:"children"`
}
