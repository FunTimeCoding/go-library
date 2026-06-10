package service

import (
	"github.com/funtimecoding/go-library/pkg/proxmox"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/inventory"
	"sync"
)

type Service struct {
	inventory *inventory.Inventory
	clients   map[string]*proxmox.Client
	sessions  sync.Map
	mutex     sync.Mutex
}
