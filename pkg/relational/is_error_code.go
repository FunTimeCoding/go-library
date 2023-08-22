package relational

import "github.com/jackc/pgconn"

func IsErrorCode(
	e error,
	code string,
) bool {
	if f, yes := e.(*pgconn.PgError); yes {
		return f.Code == code
	}

	return false
}
