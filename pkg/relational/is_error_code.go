package relational

import (
	"errors"
	"github.com/jackc/pgconn"
)

func IsErrorCode(
	e error,
	code string,
) bool {
	var f *pgconn.PgError

	if errors.As(e, &f) {
		return f.Code == code
	}

	return false
}
