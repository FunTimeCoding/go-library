package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/errors/sentry/recovery"
	"github.com/funtimecoding/go-library/pkg/log/logger"
	"github.com/funtimecoding/go-library/pkg/raid"
	"github.com/getsentry/sentry-go"
	"gorm.io/gorm"
)

func New(
	m *gorm.DB,
	logCachePath string,
	elitePath string,
	l *logger.Logger,
	h *sentry.Hub,
) *Store {
	errors.PanicOnError(m.AutoMigrate(raid.NewRaid()))
	errors.PanicOnError(m.AutoMigrate(raid.NewFight()))
	errors.PanicOnError(m.AutoMigrate(raid.NewPlayerFightStat()))
	s := &Store{
		mapper:       m,
		logger:       l,
		recovery:     recovery.New(l, h),
		logCachePath: logCachePath,
		elitePath:    elitePath,
		stop:         make(chan struct{}),
	}
	s.syncLogCache()

	return s
}
