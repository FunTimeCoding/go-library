package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gorm.io/gorm"
)

func New(m *gorm.DB) *Store {
	errors.PanicOnError(m.AutoMigrate(NewPlaybookRun()))

	return &Store{mapper: m}
}
