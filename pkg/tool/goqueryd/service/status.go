package service

import "github.com/funtimecoding/go-library/pkg/tool/goqueryd/store/status"

func (s *Service) Status() (*status.Status, error) {
	return s.store.Status()
}
