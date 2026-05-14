package argument

type ResetMachine struct {
	VMID int    `json:"vmid"`
	Node string `json:"node"`
}
