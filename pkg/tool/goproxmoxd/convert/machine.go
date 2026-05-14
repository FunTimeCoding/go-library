package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func Machine(v proxmox.VirtualMachine) server.Machine {
	vmid := int64(v.VMID)
	cpu := v.CPU
	cpus := v.CPUs
	mem := int64(v.Mem)
	maxMem := int64(v.MaxMem)
	uptime := int64(v.Uptime)
	result := server.Machine{
		Vmid:   vmid,
		Name:   v.Name,
		Node:   &v.Node,
		Status: &v.Status,
		Cpu:    &cpu,
		Cpus:   &cpus,
		Mem:    &mem,
		MaxMem: &maxMem,
		Uptime: &uptime,
	}

	if v.Tags != "" {
		result.Tags = &v.Tags
	}

	return result
}
