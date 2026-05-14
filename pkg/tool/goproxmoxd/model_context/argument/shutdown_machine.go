package argument

type ShutdownMachine struct {
	VMID int    `json:"vmid"`
	Node string `json:"node"`
}
