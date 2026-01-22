package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/maintenance_log/entry"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func New() *Store {
	database, e := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(database.AutoMigrate(&entry.Entry{}))

	return &Store{database: database}
}
