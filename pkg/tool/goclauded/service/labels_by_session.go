package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/label"

func (s *Service) LabelsBySession(sessionIdentifier string) []label.Label {
	return s.store.LabelsBySession(sessionIdentifier)
}
