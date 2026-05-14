package argument

type CreateSnapshot struct {
	VMID int    `json:"vmid"`
	Node string `json:"node"`
	Name string `json:"name"`
}
