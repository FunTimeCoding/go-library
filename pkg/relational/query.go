package relational

import (
	"errors"
	"github.com/jackc/pgx/v5"
)

func (d *Database) Query(
	sql string,
	a ...any,
) pgx.Rows {
	result, e := d.client.Query(d.context, sql, a...)

	if errors.Is(e, pgx.ErrNoRows) {
		return result
	} else if e != nil {
		panic(e)
	}

	return result
}
