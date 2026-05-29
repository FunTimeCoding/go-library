package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gokubernetesd/store/mute_rule"
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func New(path string) *Store {
	d, e := gorm.Open(sqlite.Open(path), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(d.Exec("PRAGMA journal_mode = WAL").Error)
	errors.PanicOnError(d.AutoMigrate(mute_rule.Stub()))

	return &Store{database: d}
}
