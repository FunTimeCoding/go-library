package rotation

import "github.com/funtimecoding/go-library/pkg/face"

func ContainsUser(
	r *Rotation,
	user string,
	f face.StringAlias,
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
