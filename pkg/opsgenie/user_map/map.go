package user_map

import "github.com/funtimecoding/go-library/pkg/opsgenie/user"

type Map struct {
	Users   []*user.User
	UserMap map[string]*user.User
}
