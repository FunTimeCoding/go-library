package convert

import (
	"github.com/funtimecoding/go-library/pkg/proxmox/node_status"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
)

func NodeStatus(s *node_status.Status) server.NodeStatus {
	cpu := s.Processor
	uptime := s.Uptime
	kernelVersion := s.KernelVersion
	proxmoxVersion := s.ProxmoxVersion
	memFree := s.Memory.Free
	memTotal := s.Memory.Total
	memUsed := s.Memory.Used
	swapFree := s.Swap.Free
	swapTotal := s.Swap.Total
	swapUsed := int64(s.Swap.Used)
	rootFree := s.RootFilesystem.Free
	rootTotal := s.RootFilesystem.Total
	rootUsed := s.RootFilesystem.Used

	return server.NodeStatus{
		Cpu:            &cpu,
		Uptime:         &uptime,
		KernelVersion:  &kernelVersion,
		ProxmoxVersion: &proxmoxVersion,
		LoadAverage:    &s.LoadAverage,
		Memory: &server.MemoryInfo{
			Free:  &memFree,
			Total: &memTotal,
			Used:  &memUsed,
		},
		Swap: &server.MemoryInfo{
			Free:  &swapFree,
			Total: &swapTotal,
			Used:  &swapUsed,
		},
		RootFilesystem: &server.MemoryInfo{
			Free:  &rootFree,
			Total: &rootTotal,
			Used:  &rootUsed,
		},
	}
}
