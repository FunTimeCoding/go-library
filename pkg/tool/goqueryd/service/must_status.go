package service

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/store/status"

func (s *Service) MustStatus() *status.Status {
	return s.store.MustStatus()
}
