package relational

import "github.com/funtimecoding/go-library/errors"

func (d *Database) Delete(a any) {
	errors.PanicOnError(d.mapper.Delete(a).Error)
}
