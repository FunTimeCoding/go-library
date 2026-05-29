package service

import "github.com/funtimecoding/go-library/pkg/tool/goprometheusd/inventory"

func (s *Service) Instances() []inventory.Instance {
	return s.inventory.Instances
}
