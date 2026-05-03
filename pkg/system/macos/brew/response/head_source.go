package response

type HeadSource struct {
	Locator string  `json:"url"`
	Branch  *string `json:"branch"`
	Using   *string `json:"using"`
}
