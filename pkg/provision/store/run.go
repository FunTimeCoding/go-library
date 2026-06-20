package store

import "time"

type Run struct {
	ID                  uint `gorm:"primaryKey"`
	CreatedAt           time.Time
	Scope               string
	TriggerSource       string
	DurationMillisecond int64
	Status              string
	Output              string
	ErrorOutput         string
	GitHead             string
}
