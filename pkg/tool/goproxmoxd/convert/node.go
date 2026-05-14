package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func Node(n proxmox.NodeStatus) server.Node {
	cpu := n.CPU
	maxCpu := n.MaxCPU
	mem := int64(n.Mem)
	maxMem := int64(n.MaxMem)
	disk := int64(n.Disk)
	maxDisk := int64(n.MaxDisk)
	uptime := int64(n.Uptime)

	return server.Node{
		Name:    n.Node,
		Status:  n.Status,
		Cpu:     &cpu,
		MaxCpu:  &maxCpu,
		Mem:     &mem,
		MaxMem:  &maxMem,
		Disk:    &disk,
		MaxDisk: &maxDisk,
		Uptime:  &uptime,
	}
}
