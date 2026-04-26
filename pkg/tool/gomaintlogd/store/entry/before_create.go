package entry

import (
	"gorm.io/gorm"
	"time"
)

func (e *Entry) BeforeCreate(_ *gorm.DB) error {
	if e.Timestamp.IsZero() {
		e.Timestamp = time.Now()
	}

	return nil
}
