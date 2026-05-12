package argument

import "github.com/funtimecoding/go-library/pkg/errors"

func (i *Instance) GetBoolean(name string) bool {
	v, e := i.flags.GetBool(name)
	errors.PanicOnError(e)

	return v
}
