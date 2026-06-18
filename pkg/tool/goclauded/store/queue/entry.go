package queue

import "time"

type Entry struct {
	Identifier uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	Callsign   string    `gorm:"column:callsign"`
	Kind       string    `gorm:"column:kind"`
	Body       string    `gorm:"column:body"`
	Consumed   bool      `gorm:"column:consumed"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}
