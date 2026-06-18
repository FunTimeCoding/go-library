package service

import "github.com/funtimecoding/go-library/pkg/tool/gosourced/inventory"

func New(i *inventory.Inventory) *Service {
	return &Service{inventory: i}
}
