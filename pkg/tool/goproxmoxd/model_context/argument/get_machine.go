package argument

type GetMachine struct {
	VMID int    `json:"vmid"`
	Node string `json:"node"`
}
