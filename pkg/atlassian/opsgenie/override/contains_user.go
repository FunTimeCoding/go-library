package override

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/compact"

func ContainsUser(
	o *Override,
	user string,
) bool {
	return compact.Username(o.User) == user
}
