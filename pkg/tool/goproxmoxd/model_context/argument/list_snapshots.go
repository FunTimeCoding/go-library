package argument

type ListSnapshots struct {
	VMID int    `json:"vmid"`
	Node string `json:"node"`
}
