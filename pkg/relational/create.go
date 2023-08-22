package relational

import (
	"github.com/funtimecoding/go-library/pkg/errors"
)

func (d *Database) Create(a any) {
	errors.PanicOnError(d.mapper.Create(a).Error)
}
