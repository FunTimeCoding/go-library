package service

import "github.com/funtimecoding/go-library/pkg/tool/goproxmoxd/inventory"

func (s *Service) Instances() []inventory.Instance {
	return s.inventory.Instances
}
