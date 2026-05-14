package argument

type StopMachine struct {
	VMID int    `json:"vmid"`
	Node string `json:"node"`
}
