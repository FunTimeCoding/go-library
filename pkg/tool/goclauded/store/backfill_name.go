package store

import "gorm.io/gorm"

func backfillName(d *gorm.DB) {
	d.Exec("UPDATE session SET name = callsign WHERE name IS NULL OR name = ''")
}
