package service

import (
	library "github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/claude/tracker"
	"github.com/funtimecoding/go-library/pkg/generative/anthropic/site/usage_result"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/face"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/client"
	"sync"
	"time"
)

type Service struct {
	store      *store.Store
	client     face.ClaudeSource
	memory     client.Client
	indexer    library.Indexer
	notifier   face.Notifier
	reporter   library.Reporter
	clock      func() time.Time
	logger     *logger.Logger
	harbor     string
	states     map[string]*tracker.State
	statesMu   sync.Mutex
	usage      *usage_result.Result
	usageMutex sync.RWMutex
}
