package response

type Priority struct {
	Self        string `json:"self"`
	IconLocator string `json:"iconUrl"`
	Name        string `json:"name"`
	Identifier  string `json:"id"`
}
