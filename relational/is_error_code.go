package relational

import "github.com/jackc/pgconn"

func IsErrorCode(
	fail error,
	code string,
) bool {
	if e, yes := fail.(*pgconn.PgError); yes {
		return e.Code == code
	}

	return false
}
