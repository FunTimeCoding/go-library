package service

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/client"
)

func New(
	s *store.Store,
	m client.Client,
) *Service {
	return &Service{Store: s, memory: m}
}
