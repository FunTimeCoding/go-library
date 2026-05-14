package response

type Machine struct {
	VMID   uint64  `json:"vmid"`
	Name   string  `json:"name"`
	Node   string  `json:"node"`
	Status string  `json:"status"`
	CPU    float64 `json:"cpu"`
	CPUs   int     `json:"cpus"`
	Mem    uint64  `json:"mem"`
	MaxMem uint64  `json:"max_mem"`
	Uptime uint64  `json:"uptime"`
	Tags   string  `json:"tags,omitempty"`
}
