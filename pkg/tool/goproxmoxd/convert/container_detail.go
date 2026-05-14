package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func ContainerDetail(c *proxmox.Container) server.ContainerDetail {
	vmid := int64(c.VMID)
	cpus := c.CPUs
	maxMem := int64(c.MaxMem)
	maxDisk := int64(c.MaxDisk)
	maxSwap := int64(c.MaxSwap)
	uptime := int64(c.Uptime)
	result := server.ContainerDetail{
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

	if c.ContainerConfig != nil {
		if c.ContainerConfig.Hostname != "" {
			result.Hostname = &c.ContainerConfig.Hostname
		}

		if c.ContainerConfig.Description != "" {
			result.Description = &c.ContainerConfig.Description
		}

		if c.ContainerConfig.Arch != "" {
			result.Arch = &c.ContainerConfig.Arch
		}

		if c.ContainerConfig.Cores > 0 {
			result.Cores = &c.ContainerConfig.Cores
		}

		if c.ContainerConfig.Memory > 0 {
			result.Memory = &c.ContainerConfig.Memory
		}
	}

	return result
}
