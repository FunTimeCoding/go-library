package response

type ChangeHistory struct {
	Id      string       `json:"id"`
	Created string       `json:"created"`
	Items   []ChangeItem `json:"items"`
}
