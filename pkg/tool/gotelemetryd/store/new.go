package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gorm.io/gorm"
)

func New(m *gorm.DB) *Store {
	errors.PanicOnError(m.AutoMigrate(NewUsageEvent()))

	return &Store{mapper: m}
}
