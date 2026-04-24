package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewLite(path string) *Store {
	database, e := gorm.Open(sqlite.Open(path), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(database.AutoMigrate(&Entry{}))

	return &Store{database: database}
}
