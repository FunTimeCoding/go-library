package event

import "time"

type Event struct {
	Identifier        uint              `gorm:"primaryKey;autoIncrement;column:identifier"`
	SessionIdentifier string            `gorm:"column:session_identifier"`
	Kind              string            `gorm:"column:kind"`
	Actor             string            `gorm:"column:actor"`
	CreatedAt         time.Time         `gorm:"column:created_at"`
	Metadata          map[string]string `gorm:"-"`
}
