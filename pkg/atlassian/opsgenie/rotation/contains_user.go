package rotation

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/constant"

func ContainsUser(
	r *Rotation,
	user string,
	f constant.StringAlias,
) bool {
	for _, element := range r.Participants {
		if f == nil {
			if element.Username == user {
				return true
			}
		} else {
			if f(element.Username) == user {
				return true
			}
		}
	}

	return false
}
