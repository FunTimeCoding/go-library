package argument

import "github.com/funtimecoding/go-library/pkg/errors"

func (i *Instance) GetInteger64(name string) int64 {
	v, e := i.flags.GetInt64(name)
	errors.PanicOnError(e)

	return v
}
