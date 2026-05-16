package response

type Status struct {
	Self           string         `json:"self"`
	Description    string         `json:"description"`
	IconLocator    string         `json:"iconUrl"`
	Name           string         `json:"name"`
	Identifier     string         `json:"id"`
	StatusCategory StatusCategory `json:"statusCategory"`
}
