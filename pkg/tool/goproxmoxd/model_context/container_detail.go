package model_context

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/response"
	"github.com/luthermonson/go-proxmox"
)

func containerDetail(c *proxmox.Container) response.ContainerDetail {
	result := response.ContainerDetail{
		VMID:    uint64(c.VMID),
		Name:    c.Name,
		Node:    c.Node,
		Status:  c.Status,
		CPUs:    c.CPUs,
		MaxMem:  c.MaxMem,
		MaxDisk: c.MaxDisk,
		MaxSwap: c.MaxSwap,
		Uptime:  c.Uptime,
		Tags:    c.Tags,
	}

	if c.ContainerConfig != nil {
		result.Hostname = c.ContainerConfig.Hostname
		result.Description = c.ContainerConfig.Description
		result.Arch = c.ContainerConfig.Arch
		result.Cores = c.ContainerConfig.Cores
		result.Memory = c.ContainerConfig.Memory
	}

	return result
}
