package service

import (
	"github.com/funtimecoding/go-library/pkg/prometheus"
	"github.com/funtimecoding/go-library/pkg/tool/goprometheusd/inventory"
)

func New(i *inventory.Inventory) *Service {
	return &Service{
		inventory: i,
		clients:   make(map[string]*prometheus.Client),
	}
}
