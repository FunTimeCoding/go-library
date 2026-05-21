package store

import "gorm.io/gorm"

func migrateNameToCallsign(d *gorm.DB) {
	var count int64
	d.Raw("SELECT COUNT(*) FROM pragma_table_info('session') WHERE name = 'callsign'").Scan(&count)

	if count > 0 {
		return
	}

	d.Exec("ALTER TABLE session RENAME COLUMN name TO callsign")
	d.Exec("DROP INDEX IF EXISTS idx_session_name")
}
