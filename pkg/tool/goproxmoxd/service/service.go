package service

import (
	"github.com/funtimecoding/go-library/pkg/proxmox"
	"github.com/funtimecoding/go-library/pkg/ssh"
	"github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/inventory"
	"sync"
)

type Service struct {
	inventory  *inventory.Inventory
	clients    map[string]*proxmox.Client
	sshClients map[string]*ssh.Client
	sessions   sync.Map
	mutex      sync.Mutex
}
