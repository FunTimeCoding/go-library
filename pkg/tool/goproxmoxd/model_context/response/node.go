package response

type Node struct {
	Name    string  `json:"name"`
	Status  string  `json:"status"`
	CPU     float64 `json:"cpu"`
	MaxCPU  int     `json:"max_cpu"`
	Mem     uint64  `json:"mem"`
	MaxMem  uint64  `json:"max_mem"`
	Disk    uint64  `json:"disk"`
	MaxDisk uint64  `json:"max_disk"`
	Uptime  uint64  `json:"uptime"`
}
