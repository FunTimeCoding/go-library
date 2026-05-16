package response

type ChangeHistory struct {
	Identifier string       `json:"id"`
	Created    string       `json:"created"`
	Items      []ChangeItem `json:"items"`
}
