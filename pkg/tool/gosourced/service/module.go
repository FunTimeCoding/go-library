package service

import "github.com/funtimecoding/go-library/pkg/tool/gosourced/inventory"

func (s *Service) Module(name string) (*inventory.Module, bool) {
	for _, m := range s.inventory.Modules {
		if m.Name == name {
			return &m, true
		}
	}

	return nil, false
}
