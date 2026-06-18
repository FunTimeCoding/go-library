package service

import "github.com/funtimecoding/go-library/pkg/tool/goclauded/store/queue"

func (s *Service) DrainQueue(callsign string) ([]queue.Entry, error) {
	return s.store.DrainQueue(callsign)
}
