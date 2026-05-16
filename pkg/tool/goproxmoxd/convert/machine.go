package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func Machine(v proxmox.VirtualMachine) *server.Machine {
	result := &server.Machine{
		Vmid:   int64(v.VMID),
		Name:   v.Name,
		Node:   &v.Node,
		Status: &v.Status,
		Cpu:    new(v.CPU),
		Cpus:   new(v.CPUs),
		Mem:    new(int64(v.Mem)),
		MaxMem: new(int64(v.MaxMem)),
		Uptime: new(int64(v.Uptime)),
	}

	if v.Tags != "" {
		result.Tags = &v.Tags
	}

	return result
}
