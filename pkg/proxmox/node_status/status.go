package node_status

type Swap struct {
	Free  int64 `json:"free"`
	Total int64 `json:"total"`
	Used  int   `json:"used"`
}

type BootDetail struct {
	Mode       string `json:"mode"`
	SecureBoot int    `json:"secureboot"`
}

type SamePageMerge struct {
	Shared int `json:"shared"`
}

type Kernel struct {
	Machine string `json:"machine"`
	Release string `json:"release"`
	SysName string `json:"sysname"`
	Version string `json:"version"`
}

type Memory struct {
	Available int64 `json:"available"`
	Free      int64 `json:"free"`
	Total     int64 `json:"total"`
	Used      int64 `json:"used"`
}

type ProcessorDetail struct {
	Cores     int    `json:"cores"`
	Count     int    `json:"cpus"`
	Flag      string `json:"flags"`
	Hardware  string `json:"hvm"`
	Megahertz string `json:"mhz"`
	Model     string `json:"model"`
	Sockets   int    `json:"sockets"`
	UserHertz int    `json:"user_hz"`
}

type RootFilesystem struct {
	Available int64 `json:"avail"`
	Free      int64 `json:"free"`
	Total     int64 `json:"total"`
	Used      int64 `json:"used"`
}

type Status struct {
	BootDetail      BootDetail      `json:"boot-info"`
	Idle            int             `json:"idle"`
	Kernel          Kernel          `json:"current-kernel"`
	KernelVersion   string          `json:"kversion"`
	LoadAverage     []string        `json:"loadavg"`
	Memory          Memory          `json:"memory"`
	Processor       float64         `json:"cpu"`
	ProcessorDetail ProcessorDetail `json:"cpuinfo"`
	ProxmoxVersion  string          `json:"pveversion"`
	RootFilesystem  RootFilesystem  `json:"rootfs"`
	SamePageMerge   SamePageMerge   `json:"ksm"`
	Swap            Swap            `json:"swap"`
	Uptime          int             `json:"uptime"`
	Wait            int             `json:"wait"`
}
