package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/completion"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/event"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/message"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/pool_name"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/session"
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/store/summary"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"time"
)

func New(
	path string,
	clock func() time.Time,
) *Store {
	d, e := gorm.Open(sqlite.Open(path), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(d.Exec("PRAGMA journal_mode = WAL").Error)
	errors.PanicOnError(
		d.AutoMigrate(
			session.Stub(),
			message.Stub(),
			event.Stub(),
			completion.Stub(),
			summary.Stub(),
			pool_name.Stub(),
		),
	)

	return &Store{database: d, clock: clock}
}
