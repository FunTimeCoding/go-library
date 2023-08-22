package relational

import "github.com/jackc/pgx/v5/pgxpool"

func (d *Database) Client() *pgxpool.Pool {
	return d.client
}
