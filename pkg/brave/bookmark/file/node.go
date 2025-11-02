package file

type Node struct {
	DateAdded    string  `json:"date_added"`
	DateLastUsed string  `json:"date_last_used"`
	Guid         string  `json:"guid"`
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Type         string  `json:"type"`
	Url          string  `json:"url"`
	Children     []*Node `json:"children"`
}
