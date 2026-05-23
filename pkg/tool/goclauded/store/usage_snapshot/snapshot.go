package usage_snapshot

import "time"

type Snapshot struct {
	Identifier     uint      `gorm:"primaryKey;autoIncrement;column:identifier"`
	SessionPercent int       `gorm:"column:session_percent"`
	WeeklyPercent  int       `gorm:"column:weekly_percent"`
	SonnetPercent  int       `gorm:"column:sonnet_percent"`
	CreditPercent  int       `gorm:"column:credit_percent"`
	CreatedAt      time.Time `gorm:"column:created_at"`
}
