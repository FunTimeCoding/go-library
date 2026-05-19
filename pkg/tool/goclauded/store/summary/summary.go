package summary

import "time"

type Summary struct {
	Identifier        uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	SessionIdentifier string    `gorm:"column:session_identifier"`
	Name              string    `gorm:"column:name"`
	Body              string    `gorm:"column:body"`
	CreatedAt         time.Time `gorm:"column:created_at"`
}
