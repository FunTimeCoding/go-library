package argument

type DeleteMachine struct {
	Identifier int    `json:"identifier"`
	Node       string `json:"node"`
	Purge      bool   `json:"purge"`
}
