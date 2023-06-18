package relational

import "github.com/jackc/pgx/v5"

func (d *Database) Query(
	sql string,
	arguments ...any,
) pgx.Rows {
	rows, e := d.client.Query(d.context, sql, arguments...)

	if e == pgx.ErrNoRows {
		return rows
	} else if e != nil {
		panic(e)
	}

	return rows
}
