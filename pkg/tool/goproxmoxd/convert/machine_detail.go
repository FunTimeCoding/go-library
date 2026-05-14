package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func MachineDetail(v *proxmox.VirtualMachine) server.MachineDetail {
	vmid := int64(v.VMID)
	cpu := v.CPU
	cpus := v.CPUs
	mem := int64(v.Mem)
	maxMem := int64(v.MaxMem)
	disk := int64(v.Disk)
	maxDisk := int64(v.MaxDisk)
	uptime := int64(v.Uptime)
	result := server.MachineDetail{
		Vmid:    vmid,
		Name:    v.Name,
		Node:    &v.Node,
		Status:  &v.Status,
		Cpu:     &cpu,
		Cpus:    &cpus,
		Mem:     &mem,
		MaxMem:  &maxMem,
		Disk:    &disk,
		MaxDisk: &maxDisk,
		Uptime:  &uptime,
	}

	if v.Tags != "" {
		result.Tags = &v.Tags
	}

	if v.VirtualMachineConfig != nil {
		if v.VirtualMachineConfig.Description != "" {
			result.Description = &v.VirtualMachineConfig.Description
		}

		if v.VirtualMachineConfig.OSType != "" {
			result.OsType = &v.VirtualMachineConfig.OSType
		}

		if v.VirtualMachineConfig.Sockets > 0 {
			result.Sockets = &v.VirtualMachineConfig.Sockets
		}

		if v.VirtualMachineConfig.Cores > 0 {
			result.Cores = &v.VirtualMachineConfig.Cores
		}
	}

	return result
}
