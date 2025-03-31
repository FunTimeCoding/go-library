package user

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"

type User struct {
	Name string
	Raw  *response.User
}
