package relational

import "github.com/funtimecoding/go-library/pkg/errors"

func (d *Database) UserExists(n string) bool {
	var result bool
	errors.PanicOnError(
		d.QueryRow(
			"SELECT EXISTS(SELECT 1 FROM pg_roles WHERE rolname = $1)",
			n,
		).Scan(&result),
	)

	return result
}
