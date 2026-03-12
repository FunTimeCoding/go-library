package node_status

type Kernel struct {
	Machine string `json:"machine"`
	Release string `json:"release"`
	SysName string `json:"sysname"`
	Version string `json:"version"`
}
