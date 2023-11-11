package relational

import "github.com/funtimecoding/go-library/pkg/errors"

func (d *Database) Execute(
	sql string,
	arguments ...any,
) {
	_, e := d.client.Exec(d.context, sql, arguments...)
	errors.PanicOnError(e)
}
