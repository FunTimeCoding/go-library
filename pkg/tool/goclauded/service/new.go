package service

import (
	library "github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/tracker"
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
	r library.Reporter,
	harbor string,
	clock func() time.Time,
	l *logger.Logger,
) *Service {
	return &Service{
		store:    s,
		client:   c,
		memory:   m,
		indexer:  i,
		notifier: n,
		reporter: r,
		harbor:   harbor,
		states:   make(map[string]*tracker.State),
		clock:    clock,
		logger:   l,
	}
}
