package store

import "gorm.io/gorm"

func migrateColumns(d *gorm.DB) {
	var count int64
	d.Raw(
		"SELECT COUNT(*) FROM pragma_table_info('session') WHERE name = 'cwd'",
	).Scan(&count)

	if count > 0 {
		d.Exec("ALTER TABLE session RENAME COLUMN cwd TO work_directory")
	}
}
