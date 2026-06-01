package service

import (
	"github.com/funtimecoding/go-library/pkg/prometheus/alertmanager"
	"github.com/funtimecoding/go-library/pkg/tool/goalertmanagerd/inventory"
)

func New(i *inventory.Inventory) *Service {
	return &Service{
		inventory: i,
		clients:   make(map[string]*alertmanager.Client),
	}
}
