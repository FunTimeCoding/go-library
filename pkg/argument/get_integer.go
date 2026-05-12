package argument

import "github.com/funtimecoding/go-library/pkg/errors"

func (i *Instance) GetInteger(name string) int {
	v, e := i.flags.GetInt(name)
	errors.PanicOnError(e)

	return v
}
