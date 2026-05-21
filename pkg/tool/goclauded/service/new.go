package service

import (
	library "github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/face"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/client"
	"time"
)

func New(
	s *store.Store,
	c face.ClaudeSource,
	m client.Client,
	i library.Indexer,
	n face.Notifier,
	clock func() time.Time,
	l *logger.Logger,
) *Service {
	return &Service{
		store:    s,
		client:   c,
		memory:   m,
		indexer:  i,
		notifier: n,
		clock:    clock,
		logger:   l,
	}
}
