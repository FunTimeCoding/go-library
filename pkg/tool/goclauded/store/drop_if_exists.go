package store

import (
	"fmt"
	"github.com/funtimecoding/go-library/pkg/errors"
	"gorm.io/gorm"
)

func dropIfExists(
	d *gorm.DB,
	table string,
	column string,
) {
	if !columnExists(d, table, column) {
		return
	}

	errors.PanicOnError(
		d.Exec(
			fmt.Sprintf(
				"ALTER TABLE %s DROP COLUMN %s",
				table,
				column,
			),
		).Error,
	)
}
