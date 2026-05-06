package store

import "time"

type HighstateRun struct {
	ID                  uint `gorm:"primaryKey"`
	CreatedAt           time.Time
	TriggerSource       string
	DurationMillisecond int64
	Status              string
	Output              string
	ErrorOutput         string
	GitHead             string
}
