package argument

import "github.com/funtimecoding/go-library/pkg/strings/split"

func (i *Instance) Slice(name string) []string {
	v := i.GetString(name)

	if v == "" {
		return []string{}
	}

	return split.Comma(v)
}
