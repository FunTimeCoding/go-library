package user_map

import "github.com/funtimecoding/go-library/pkg/opsgenie/user"

func New(v []*user.User) *Map {
	m := make(map[string]*user.User, len(v))

	for _, element := range v {
		m[element.Identifier] = element
	}

	return &Map{Users: v, UserMap: m}
}
