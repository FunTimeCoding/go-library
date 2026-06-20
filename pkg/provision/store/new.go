package store

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"gorm.io/gorm"
)

func New(m *gorm.DB, tableName string) *Store {
	errors.PanicOnError(m.Table(tableName).AutoMigrate(&Run{}))

	return &Store{mapper: m, tableName: tableName}
}
