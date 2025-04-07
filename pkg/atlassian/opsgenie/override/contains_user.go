package override

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"

func ContainsUser(
	o *Override,
	user string,
	f constant.StringAlias,
) bool {
	if f == nil {
		return o.User == user
	}

	return f(o.User) == user
}
