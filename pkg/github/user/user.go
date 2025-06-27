package user

import "github.com/google/go-github/v70/github"

type User struct {
	Name string
	Raw  *github.User
}
