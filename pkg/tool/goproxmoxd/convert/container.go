package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func Container(c proxmox.Container) server.Container {
	vmid := int64(c.VMID)
	cpus := c.CPUs
	maxMem := int64(c.MaxMem)
	maxDisk := int64(c.MaxDisk)
	maxSwap := int64(c.MaxSwap)
	uptime := int64(c.Uptime)
	result := server.Container{
		Vmid:    vmid,
		Name:    c.Name,
		Node:    &c.Node,
		Status:  &c.Status,
		Cpus:    &cpus,
		MaxMem:  &maxMem,
		MaxDisk: &maxDisk,
		MaxSwap: &maxSwap,
		Uptime:  &uptime,
	}

	if c.Tags != "" {
		result.Tags = &c.Tags
	}

	return result
}
