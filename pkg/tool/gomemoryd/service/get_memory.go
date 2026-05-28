package service

import "github.com/funtimecoding/go-library/pkg/tool/gomemoryd/store"

func (s *Service) GetMemory(identifier int64) (*store.Memory, error) {
	return s.store.GetMemory(identifier)
}
