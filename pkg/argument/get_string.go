package argument

import "github.com/funtimecoding/go-library/pkg/errors"

func (i *Instance) GetString(name string) string {
	v, e := i.flags.GetString(name)
	errors.PanicOnError(e)

	return v
}
