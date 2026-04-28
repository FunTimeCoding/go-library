package store

import (
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"gorm.io/gorm"
)

type Store struct {
	mapper            *gorm.DB
	logger            *logger.Logger
	recovery          *recovery.Recovery
	logCachePath      string
	eliteInsightsPath string
	stop              chan struct{}
}
