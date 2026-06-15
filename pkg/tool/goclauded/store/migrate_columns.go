package store

import (
	"github.com/funtimecoding/go-library/pkg/tool/goclauded/constant"
	"gorm.io/gorm"
)

func migrateColumns(d *gorm.DB) {
	renameIfExists(d, "session", "cwd", "work_directory")
	renameIfExists(d, "session", "created_at", "started_at")
	renameIfExists(d, "session", "updated_at", "last_active_at")
	renameIfExists(d, "event", constant.Target, "scope")
	renameIfExists(d, "event", constant.SessionName, "actor")
}
