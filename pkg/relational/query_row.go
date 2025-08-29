package relational

import "github.com/jackc/pgx/v5"

func (d *Database) QueryRow(
	sql string,
	a ...any,
) pgx.Row {
	return d.client.QueryRow(d.context, sql, a...)
}
