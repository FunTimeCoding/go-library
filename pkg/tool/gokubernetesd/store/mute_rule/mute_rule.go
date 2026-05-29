package mute_rule

import "time"

type MuteRule struct {
	Identifier uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	Reason     string    `gorm:"column:reason"`
	Message    string    `gorm:"column:message"`
	Cluster    string    `gorm:"column:cluster"`
	CreatedAt  time.Time `gorm:"column:created_at"`
}
