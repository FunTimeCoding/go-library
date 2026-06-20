package create_machine

import "github.com/luthermonson/go-proxmox"

func option(
	name string,
	value any,
) proxmox.VirtualMachineOption {
	return proxmox.VirtualMachineOption{Name: name, Value: value}
}
