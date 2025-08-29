package relational

import "github.com/funtimecoding/go-library/pkg/errors"

func (d *Database) DeleteWhere(
	t any,
	query string,
	a ...any,
) {
	errors.PanicOnError(d.Mapper().Where(query, a...).Delete(t).Error)
}
