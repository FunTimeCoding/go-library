package service

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/inventory"
	"sync"
)

type Service struct {
	inventory *inventory.Inventory
	clients   map[string]*alertmanager.Client
	sessions  sync.Map
	mu        sync.Mutex
}
