package user

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic/response"

type User struct {
	Name string
	Raw  *response.User
}
