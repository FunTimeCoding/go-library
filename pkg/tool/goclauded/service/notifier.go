package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/face"

func (s *Service) Notifier() face.Notifier {
	return s.notifier
}
