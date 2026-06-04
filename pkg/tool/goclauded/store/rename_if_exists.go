package store

import (
	"fmt"
	"gorm.io/gorm"
)

func renameIfExists(
	d *gorm.DB,
	table string,
	from string,
	to string,
) {
	var count int64
	d.Raw(
		"SELECT COUNT(*) FROM pragma_table_info(?) WHERE name = ?",
		table,
		from,
	).Scan(&count)

	if count > 0 {
		d.Exec(
			fmt.Sprintf(
				"ALTER TABLE %s RENAME COLUMN %s TO %s",
				table,
				from,
				to,
			),
		)
	}
}
