package store

import (
	"gorm.io/gorm"
	"time"
)

type Store struct {
	database *gorm.DB
	clock    func() time.Time
}
