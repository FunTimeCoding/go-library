package relational

import "github.com/funtimecoding/go-library/pkg/errors"

func (d *Database) Execute(
	sql string,
	a ...any,
) {
	_, e := d.client.Exec(d.context, sql, a...)
	errors.PanicOnError(e)
}
