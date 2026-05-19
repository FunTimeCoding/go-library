package event

import "time"

type Event struct {
	Identifier        uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	SessionIdentifier string    `gorm:"column:session_identifier"`
	Kind              string    `gorm:"column:kind"`
	Name              string    `gorm:"column:name"`
	Target            string    `gorm:"column:target"`
	Body              string    `gorm:"column:body"`
	CreatedAt         time.Time `gorm:"column:created_at"`
}
