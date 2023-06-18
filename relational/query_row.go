package relational

import "github.com/jackc/pgx/v5"

func (d *Database) QueryRow(
	sql string,
	arguments ...any,
) pgx.Row {
	return d.client.QueryRow(d.context, sql, arguments...)
}
