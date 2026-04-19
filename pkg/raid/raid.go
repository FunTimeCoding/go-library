package raid

import "time"

type Raid struct {
	ID        uint      `gorm:"primaryKey"`
	Name      string
	Date      time.Time
	CreatedAt time.Time
}
