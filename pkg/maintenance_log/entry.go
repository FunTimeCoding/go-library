package maintenance_log

import (
	"gorm.io/gorm"
	"time"
)

type Entry struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	Timestamp   time.Time `json:"timestamp" gorm:"not null;index"`
	Action      string    `json:"action" gorm:"not null;type:text"`
	User        string    `json:"user" gorm:"not null;index"`
	System      string    `json:"system" gorm:"index"`
	Service     string    `json:"service" gorm:"index"`
	Description string    `json:"description" gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (Entry) TableName() string {
	return "maintenance_log_entries"
}

func (e *Entry) BeforeCreate(tx *gorm.DB) error {
	if e.Timestamp.IsZero() {
		e.Timestamp = time.Now()
	}
	return nil
}
