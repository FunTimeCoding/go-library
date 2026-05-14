package store

import "time"

type UsageEvent struct {
	ID                  uint   `gorm:"primaryKey"`
	Tool                string `gorm:"not null;index"`
	Surface             string `gorm:"not null;index"`
	Actor               string `gorm:"not null;index"`
	Outcome             string `gorm:"not null"`
	DurationMillisecond int64
	Detail              *string   `gorm:"type:jsonb"`
	CreatedAt           time.Time `gorm:"index"`
}
