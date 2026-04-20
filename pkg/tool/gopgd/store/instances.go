package store

import "github.com/funtimecoding/go-library/pkg/tool/gopgd/config"

func (s *Store) Instances() []config.Instance {
	return s.configuration.Instances
}
