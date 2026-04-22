package store

import "github.com/funtimecoding/go-library/pkg/tool/gopgd/inventory"

func (s *Store) Instances() []inventory.Instance {
	return s.inventory.Instances
}
