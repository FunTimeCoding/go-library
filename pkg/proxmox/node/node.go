package node

type Node struct {
	Name           string  `json:"node"`
	Disk           int64   `json:"disk"`
	DiskLimit      int64   `json:"maxdisk"`
	Fingerprint    string  `json:"ssl_fingerprint"`
	Identifier     string  `json:"id"`
	Level          string  `json:"level"`
	Memory         int64   `json:"mem"`
	MemoryLimit    int64   `json:"maxmem"`
	Processor      float64 `json:"cpu"`
	ProcessorLimit int     `json:"maxcpu"`
	Status         string  `json:"status"`
	Type           string  `json:"type"`
	Uptime         int     `json:"uptime"`
}
