package user

import "github.com/google/go-github/v78/github"

func New(v *github.User) *User {
	return &User{Name: v.GetLogin(), Raw: v}
}
