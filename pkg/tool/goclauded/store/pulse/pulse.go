package pulse

import "time"

type Pulse struct {
	Identifier        uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	SessionIdentifier string    `gorm:"column:session_identifier"`
	FromName          string    `gorm:"column:from_name"`
	Body              string    `gorm:"column:body"`
	Consumed          bool      `gorm:"column:consumed"`
	CreatedAt         time.Time `gorm:"column:created_at"`
}
