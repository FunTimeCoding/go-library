package relational

import "github.com/funtimecoding/go-library/pkg/errors"

func (d *Database) Timezone() string {
	var result string
	errors.PanicOnError(d.QueryRow("SHOW TIMEZONE").Scan(&result))

	return result
}
