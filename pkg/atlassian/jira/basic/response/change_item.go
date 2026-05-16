package response

type ChangeItem struct {
	Field           string `json:"field"`
	FieldType       string `json:"fieldtype"`
	FieldIdentifier string `json:"fieldId"`
	From            string `json:"from"`
	FromString      string `json:"fromString"`
	To              string `json:"to"`
	ToString        string `json:"toString"`
}
