package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/funtimecoding/go-library/pkg/tool/gomaintlogd/store/entry"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewPostgres(dsn string) *Store {
	database, e := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	errors.PanicOnError(e)
	errors.PanicOnError(database.AutoMigrate(entry.New()))

	return &Store{database: database}
}
