package service

import (
	library "github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/face"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/session_cache"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/client"
	"time"
)

func New(
	s *store.Store,
	c face.ClaudeSource,
	m client.Client,
	summaryIndexer library.Indexer,
	completionIndexer library.Indexer,
	n face.Notifier,
	r library.Reporter,
	harbor string,
	clock func() time.Time,
	l *logger.Logger,
) *Service {
	return &Service{
		store:             s,
		client:            c,
		memory:            m,
		summaryIndexer:    summaryIndexer,
		completionIndexer: completionIndexer,
		notifier:          n,
		reporter:          r,
		harbor:            harbor,
		cache:             session_cache.New(),
		clock:             clock,
		logger:            l,
		lastMemoryPoll:    clock().UTC().Format("2006-01-02T15:04:05Z"),
	}
}
