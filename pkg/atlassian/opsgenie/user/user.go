package user

import "github.com/opsgenie/opsgenie-go-sdk-v2/user"

type User struct {
	Identifier string
	Name       string
	FullName   string

	Raw *user.User
}
