package relational

import "github.com/funtimecoding/go-library/pkg/errors"

func (d *Database) Save(a any) {
	errors.PanicOnError(d.mapper.Save(a).Error)
}
