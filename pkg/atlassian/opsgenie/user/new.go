package user

import "github.com/opsgenie/opsgenie-go-sdk-v2/user"

func New(v *user.User) *User {
	return &User{
		Identifier: v.Id,
		Name:       v.Username,
		FullName:   v.FullName,
		Raw:        v,
	}
}
