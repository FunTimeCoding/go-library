package service

import "github.com/funtimecoding/go-library/pkg/tool/gosourced/inventory"

func (s *Service) Modules() []inventory.Module {
	return s.inventory.Modules
}
