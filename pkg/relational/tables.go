package relational

import "github.com/funtimecoding/go-library/pkg/errors"

func (d *Database) Tables() []string {
	result, e := d.Mapper().Migrator().GetTables()
	errors.PanicOnError(e)

	return result
}
