package store

import "gorm.io/gorm"

func migrateColumns(d *gorm.DB) {
	renameIfExists(d, "session", "cwd", "work_directory")
	renameIfExists(d, "session", "created_at", "started_at")
	renameIfExists(d, "session", "updated_at", "last_active_at")
}
