package override

import "github.com/funtimecoding/go-library/pkg/face"

func ContainsUser(
	o *Override,
	user string,
	f face.StringAlias,
) bool {
	if f == nil {
		return o.User == user
	}

	return f(o.User) == user
}
