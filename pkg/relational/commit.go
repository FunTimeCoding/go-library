package relational

import (
	"github.com/funtimecoding/go-library/pkg/errors"
	"github.com/jackc/pgx/v5"
)

func (d *Database) Commit(t pgx.Tx) {
	errors.PanicOnError(t.Commit(d.context))
}
