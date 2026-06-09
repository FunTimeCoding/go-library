package service

import (
	"github.com/funtimecoding/go-library/pkg/proxmox"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/inventory"
)

func New(i *inventory.Inventory) *Service {
	return &Service{
		inventory: i,
		clients:   make(map[string]*proxmox.Client),
	}
}
