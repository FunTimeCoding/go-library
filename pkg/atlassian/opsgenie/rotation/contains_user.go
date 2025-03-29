package rotation

import "github.com/funtimecoding/go-library/pkg/atlassian/opsgenie/compact"

func ContainsUser(
	r *Rotation,
	user string,
) bool {
	for _, element := range r.Participants {
		if compact.Username(element.Username) == user {
			return true
		}
	}

	return false
}
