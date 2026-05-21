package service

import (
	library "github.com/funtimecoding/go-library/pkg/face"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/face"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store"
	"github.com/funtimecoding/go-library/pkg/tool/gomemoryd/client"
	"time"
)

type Service struct {
	store    *store.Store
	client   face.ClaudeSource
	memory   client.Client
	indexer  library.Indexer
	notifier face.Notifier
	clock    func() time.Time
	logger   *logger.Logger
}
