package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/notifier"
	"gorm.io/gorm"
	"time"
)

type Store struct {
	database *gorm.DB
	notifier *notifier.Notifier
	clock    func() time.Time
}
