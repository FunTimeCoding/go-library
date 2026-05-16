package model_context

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/model_context/machine_detail"
	"github.com/luthermonson/go-proxmox"
)

func machineDetail(v *proxmox.VirtualMachine) *machine_detail.Detail {
	result := machine_detail.New()
	result.Identifier = uint64(v.VMID)
	result.Name = v.Name
	result.Node = v.Node
	result.Status = v.Status
	result.CPU = v.CPU
	result.CPUs = v.CPUs
	result.Mem = v.Mem
	result.MaxMem = v.MaxMem
	result.Disk = v.Disk
	result.MaxDisk = v.MaxDisk
	result.Uptime = v.Uptime
	result.Tags = v.Tags

	if v.VirtualMachineConfig != nil {
		result.Description = v.VirtualMachineConfig.Description
		result.OSType = v.VirtualMachineConfig.OSType
		result.Sockets = v.VirtualMachineConfig.Sockets
		result.Cores = v.VirtualMachineConfig.Cores
	}

	return result
}
