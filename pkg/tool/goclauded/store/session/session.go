package session

import "time"

type Session struct {
	Identifier          string    `gorm:"primaryKey;column:identifier"`
	Name                string    `gorm:"column:name"`
	Callsign            *string   `gorm:"uniqueIndex;column:callsign"`
	ModelContextSession string    `gorm:"column:model_context_session"`
	Topic               string    `gorm:"column:topic"`
	Files               string    `gorm:"column:files"`
	Alias               *string   `gorm:"uniqueIndex;column:alias"`
	Description         string    `gorm:"column:description"`
	Slug                string    `gorm:"column:slug"`
	FirstMessage        string    `gorm:"column:first_message"`
	TurnCount           int       `gorm:"column:turn_count"`
	Lines               int       `gorm:"column:lines"`
	WorkDirectory       string    `gorm:"column:work_directory"`
	Branch              string    `gorm:"column:branch"`
	SessionTimestamp    string    `gorm:"column:session_timestamp"`
	Listening           bool      `gorm:"column:listening"`
	TimedOut            string    `gorm:"column:timed_out"`
	NeedsRoster         bool      `gorm:"column:needs_roster"`
	NeedsReannounce     bool      `gorm:"column:needs_reannounce"`
	LastSeen            time.Time `gorm:"column:last_seen"`
	CreatedAt           time.Time `gorm:"column:created_at"`
	UpdatedAt           time.Time `gorm:"column:updated_at"`
}
