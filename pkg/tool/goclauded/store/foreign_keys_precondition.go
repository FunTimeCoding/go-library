package store

import (
	"fmt"
	"gorm.io/gorm"
	"log"
)

func foreignKeysPrecondition(d *gorm.DB) bool {
	checks := []foreignKeyCheck{
		{"event", "session_identifier", "session"},
		{"completion", "session_identifier", "session"},
		{"summary", "session_identifier", "session"},
		{"label", "session_identifier", "session"},
		{"pulse", "session_identifier", "session"},
		{"event_metadata", "event_identifier", "event"},
	}

	for _, c := range checks {
		var count int64
		d.Raw(
			fmt.Sprintf(
				`SELECT COUNT(*) FROM %s
				WHERE %s IS NULL
				OR %s NOT IN (SELECT identifier FROM %s)`,
				c.child,
				c.column,
				c.column,
				c.parent,
			),
		).Scan(&count)

		if count > 0 {
			log.Printf(
				"foreign key precondition failed: %s has %d rows with invalid %s",
				c.child,
				count,
				c.column,
			)

			return false
		}
	}

	return true
}
