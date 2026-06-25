package argument

type CloneMachine struct {
	Identifier int    `json:"identifier"`
	Name       string `json:"name"`
	Node       string `json:"node"`
	NewID      int    `json:"new_identifier"`
	Full       bool   `json:"full"`
	Storage    string `json:"storage"`
	Snapshot   string `json:"snapshot"`
}
