package user

import "github.com/funtimecoding/go-library/pkg/atlassian/confluence/basic_client/response"

func New(r *response.User) *User {
	return &User{
		Name: r.DisplayName,
		Raw:  r,
	}
}
