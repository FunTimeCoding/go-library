package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func NewPostgres(dsn string) *Store {
	database, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(database.AutoMigrate(&Entry{}))

	return &Store{database: database}
}

func NewSQLite(path string) *Store {
	database, e := gorm.Open(sqlite.Open(path), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(database.AutoMigrate(&Entry{}))

	return &Store{database: database}
}
