package response

type Machine struct {
	Identifier uint64  `json:"identifier"`
	Name       string  `json:"name"`
	Node       string  `json:"node"`
	Status     string  `json:"status"`
	CPU        float64 `json:"cpu"`
	CPUs       int     `json:"cpus"`
	Mem        uint64  `json:"mem"`
	MaxMem     uint64  `json:"max_mem"`
	Uptime     uint64  `json:"uptime"`
	Tags       string  `json:"tags,omitempty"`
}
