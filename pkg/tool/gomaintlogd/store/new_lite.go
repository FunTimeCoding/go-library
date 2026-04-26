package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store/entry"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewLite(path string) *Store {
	database, e := gorm.Open(sqlite.Open(path), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(database.AutoMigrate(entry.New()))

	return &Store{database: database}
}
