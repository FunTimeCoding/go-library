package convert

import (
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/generated/server"
	"github.com/luthermonson/go-proxmox"
)

func Network(n proxmox.NodeNetwork) server.Network {
	active := int(n.Active) == 1
	autostart := n.Autostart == 1
	result := server.Network{
		Iface:     n.Iface,
		Active:    &active,
		Autostart: &autostart,
	}

	if n.Type != "" {
		result.Type = &n.Type
	}

	if n.Address != "" {
		result.Address = &n.Address
	}

	if n.CIDR != "" {
		result.Cidr = &n.CIDR
	}

	if n.Gateway != "" {
		result.Gateway = &n.Gateway
	}

	if n.Netmask != "" {
		result.Netmask = &n.Netmask
	}

	if n.BridgePorts != "" {
		result.BridgePorts = &n.BridgePorts
	}

	if n.Comments != "" {
		result.Comments = &n.Comments
	}

	return result
}
