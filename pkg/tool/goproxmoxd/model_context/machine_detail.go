package model_context

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/response"
	"github.com/luthermonson/go-proxmox"
)

func machineDetail(v *proxmox.VirtualMachine) response.MachineDetail {
	result := response.MachineDetail{
		VMID:    uint64(v.VMID),
		Name:    v.Name,
		Node:    v.Node,
		Status:  v.Status,
		CPU:     v.CPU,
		CPUs:    v.CPUs,
		Mem:     v.Mem,
		MaxMem:  v.MaxMem,
		Disk:    v.Disk,
		MaxDisk: v.MaxDisk,
		Uptime:  v.Uptime,
		Tags:    v.Tags,
	}

	if v.VirtualMachineConfig != nil {
		result.Description = v.VirtualMachineConfig.Description
		result.OSType = v.VirtualMachineConfig.OSType
		result.Sockets = v.VirtualMachineConfig.Sockets
		result.Cores = v.VirtualMachineConfig.Cores
	}

	return result
}
