package service

import (
	"github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/tool/gosproutd/store"
)

func New(
	s *store.Store,
	n face.EventNotifier,
) *Service {
	return &Service{
		store:    s,
		notifier: n,
	}
}
