package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(dsn string) *Store {
	database, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(database.AutoMigrate(&Entry{}))

	return &Store{database: database}
}
