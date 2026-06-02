package service

import "github.com/funtimecoding/go-library/pkg/tool/gosproutd/store/seed"

func (s *Service) Seeds() []seed.Seed {
	return s.store.Seeds()
}
