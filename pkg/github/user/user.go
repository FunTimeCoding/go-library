package user

import "github.com/google/go-github/v85/github"

type User struct {
	Name string
	Raw  *github.User
}
