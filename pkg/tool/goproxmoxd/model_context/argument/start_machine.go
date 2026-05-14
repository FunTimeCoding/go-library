package argument

type StartMachine struct {
	VMID int    `json:"vmid"`
	Node string `json:"node"`
}
