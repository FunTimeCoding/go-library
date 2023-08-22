package relational

import (
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (d *Database) DeleteWhere(
	a any,
	query string,
	arguments ...any,
) {
	errors.PanicOnError(d.Mapper().Where(query, arguments...).Delete(a).Error)
}
